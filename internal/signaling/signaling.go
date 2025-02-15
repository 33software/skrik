package signaling

import (
	"encoding/json"
	"skrik/internal/auth"
	"log"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type SignalMessage struct {
	Type string          `json:"type"`
	From json.Number     `json:"from"`
	To   json.Number     `json:"to"`
	Data json.RawMessage `json:"data"`
}

type SignalConnection struct {
	UserID int
	Conn   *websocket.Conn
}

type SignalController struct {
	connections map[int]*websocket.Conn
	register    chan *SignalConnection
	unregister  chan *SignalConnection
	signals     chan SignalMessage
}

func NewSignalController(app *fiber.App) *SignalController {
	sc := &SignalController{
		connections: make(map[int]*websocket.Conn),
		register:    make(chan *SignalConnection),
		unregister:  make(chan *SignalConnection),
		signals:     make(chan SignalMessage),
	}
	go sc.Run()
	app.Get("/api/voicews", auth.Middleware(), websocket.New(sc.WSHandler))
	return sc
}

func (sc *SignalController) WSHandler(c *websocket.Conn) {
	userID := c.Locals("userid").(int)
	conn := &SignalConnection{
		UserID: userID,
		Conn:   c,
	}
	sc.register <- conn

	for {
		var msg SignalMessage
		if err := c.ReadJSON(&msg); err != nil {
			break
		}
		if msg.From == "" {
			msg.From = json.Number(strconv.Itoa(userID))
		}
		sc.signals <- msg
	}

	sc.unregister <- conn
	c.Close()
}

func (sc *SignalController) Run() {
	for {
		select {
		case conn := <-sc.register:
			sc.connections[conn.UserID] = conn.Conn

		case msg := <-sc.signals:
			recipientID, err := msg.To.Int64()
			if err != nil {
				log.Println("invalid recipient id. err: ", err)
				continue
			}
			recipient, exists := sc.connections[int(recipientID)]
			if !exists {
				log.Printf("user %d isn't connected", int(recipientID))
				continue
			}
			if err := recipient.WriteJSON(msg); err != nil {
				log.Printf("error wtiring a message to user %d. err: %v", int(recipientID), err)
			}

		case conn := <-sc.unregister:
			delete(sc.connections, conn.UserID)
		}
	}
}