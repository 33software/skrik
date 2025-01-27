package controllers

import (
	"log"
	"skrik/internal/auth"
	"skrik/internal/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type ChatController struct {
	usecase     usecases.ChatUsecaseInterface
	connections map[string]map[uint]*websocket.Conn
	register    chan *RoomConnection
	broadcast   chan RoomMessage
	unregister  chan *RoomConnection
}

type RoomConnection struct {
	UserID uint
	RoomID string
	Conn   *websocket.Conn
}

type RoomMessage struct {
	UserID  uint
	RoomID  string
	Message string
}

func NewChatController(cu usecases.ChatUsecaseInterface, app *fiber.App) {
	handler := &ChatController{
		usecase:     cu,
		connections: make(map[string]map[uint]*websocket.Conn),
		register:    make(chan *RoomConnection),
		broadcast:   make(chan RoomMessage),
		unregister:  make(chan *RoomConnection),
	}
	go handler.Run()
	app.Get("/ws", auth.Middleware(), websocket.New(handler.WSHandler))
	app.Get("/gethistory", auth.Middleware(), handler.GetMessages)

}
func (cc *ChatController) WSHandler(c *websocket.Conn) {
	useridtemp := c.Locals("userid").(int)
	userID := uint(useridtemp)
	roomID := c.Query("roomid")
	if roomID == "" {
		c.Close()
		return
	}
	cc.register <- &RoomConnection{UserID: userID, RoomID: roomID, Conn: c}

	for {
		var msg string
		err := c.ReadJSON(&msg)
		if err != nil {
			c.Close()
			break
		}
		if err := cc.usecase.SaveMessage(roomID, userID, msg); err != nil {
			continue
		}
		cc.broadcast <- RoomMessage{UserID: userID, RoomID: roomID, Message: msg}
	}
	cc.unregister <- &RoomConnection{UserID: userID, RoomID: roomID, Conn: c}
}

func (cc *ChatController) Run() {
	for {
		select {
		case conn := <-cc.register:
			_, err := cc.usecase.CreateRoomIfNotExists(conn.RoomID)
			if err != nil {
				log.Println("err: ", err)
			}
			if cc.connections[conn.RoomID] == nil {
				cc.connections[conn.RoomID] = make(map[uint]*websocket.Conn)
			}
			cc.connections[conn.RoomID][conn.UserID] = conn.Conn

		case msg := <-cc.broadcast:
			for user, conn := range cc.connections[msg.RoomID] {
				if user == msg.UserID {
					continue
				}
				err := conn.WriteJSON(fiber.Map{"sender": msg.UserID, "message": msg.Message})
				if err != nil {
					log.Printf("error occured while broadcasting to user %d in room %s. err: %v", user, msg.RoomID, err)
				}
			}
		case conn := <-cc.unregister:
			if cc.connections[conn.RoomID] != nil {
				delete(cc.connections[conn.RoomID], conn.UserID)
			}
		}
	}

}

func (cc *ChatController) GetMessages(c *fiber.Ctx) error {
	roomIDstring := c.FormValue("roomID")
	roomID, err := strconv.Atoi(roomIDstring)
	if err != nil {
		return err
	}
	messages, err := cc.usecase.GetMessagesByRoomID(uint(roomID))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(messages)

}
