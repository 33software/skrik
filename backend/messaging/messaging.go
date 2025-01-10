package messaging

import (
	"encoding/json"
	"log"
	jwtGen "skrik/JWT"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v4"
)

type wrap struct {
	conn *websocket.Conn
	mu   sync.Mutex
}
type connections struct {
	userid map[int]*wrap
	mu     sync.Mutex
}

var connManager = connections{
	userid: make(map[int]*wrap),
}

type Room struct {
	id    int
	msgCh chan []byte
	users map[int]*wrap
	mu    sync.Mutex
}
type RoomManager struct {
	Rooms map[int]*Room
	mu    sync.Mutex
}

var roomManager = &RoomManager{
	Rooms: make(map[int]*Room),
}

type Message struct {
	Messagetype string
	Sender      int
	Content     string
}

func (rm *RoomManager) CreateRoom(id int) *Room {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	_, exists := rm.Rooms[id]
	if exists {
		log.Printf("the room with id %d already exists, aborting", id)
		return nil
	}
	room := &Room{
		id:    id,
		users: make(map[int]*wrap),
		msgCh: make(chan []byte, 100),
	}
	rm.Rooms[id] = room
	go room.StartProcessing()
	return room
}
func (rm *RoomManager) DeleteRoom(id int) {
	rm.mu.Lock()
	room := rm.Rooms[id]
	close(room.msgCh)
	delete(rm.Rooms, id)
	rm.mu.Unlock()
}
func (room *Room) AddUser(userid int, user *wrap) {
	_, exists := room.users[userid]
	if exists {
		log.Printf("user %d has already joined the room", userid)
		return
	}
	room.users[userid] = user

}
func (room *Room) DeleteUser(userid int) {
	room.mu.Lock()
	delete(room.users, userid)
	if len(room.users) == 0 {
		close(room.msgCh)
	}
	room.mu.Unlock()
}
func (room *Room) StartProcessing() {
	log.Println("processing started")
	for msg := range room.msgCh {
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("couldn't deserialize message", err)
		}
		room.mu.Lock()
		for userid, user := range room.users {
			if userid == message.Sender {
				continue
			}
			user.mu.Lock()
			err := user.conn.WriteMessage(websocket.TextMessage, msg)
			user.mu.Unlock()
			if err != nil {
				connManager.mu.Lock()
				delete(connManager.userid, userid)
				connManager.mu.Unlock()
				delete(room.users, userid)
			}
		}
		room.mu.Unlock()
	}
}
func Test(app *fiber.App) {
	app.Get("room/:id", jwtGen.JwtProtected(), websocket.New(wsHandler))
}
func wsHandler(c *websocket.Conn) {
	roomid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println("couldn't conver room id to int")
		return
	}
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	useridfloat := claims["userid"].(float64)
	userid := int(useridfloat)
	connManager.mu.Lock()
	connManager.userid[userid] = &wrap{conn: c}
	connManager.mu.Unlock()
	room, exists := roomManager.Rooms[roomid]
	if !exists {
		room = roomManager.CreateRoom(roomid)
	}
	room.mu.Lock()
	room.AddUser(userid, connManager.userid[userid])
	room.mu.Unlock()
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("lost conection")
			c.Close()
			connManager.mu.Lock()
			delete(connManager.userid, userid)
			connManager.mu.Unlock()
			room.DeleteUser(userid)
			break
		}
		message, err := json.Marshal(Message{Messagetype: "text", Sender: userid, Content: string(msg)})
		if err != nil {
			log.Println("couldn't serialize message")
			continue
		}
		room.msgCh <- message
	}
}
