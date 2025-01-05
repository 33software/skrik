package signaling

import (
	"encoding/json"
	"log"
	jwtGen "skrik/JWT"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v4"
)

type SignalMessage struct {
	Type string          `json:"type"`
	From json.Number     `json:"from"`
	To   json.Number     `json:"to"`
	Data json.RawMessage `json:"data"`
}

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

func VoiceHandler(app *fiber.App) {
	app.Get("/api/voicews", jwtGen.JwtProtected(), websocket.New(signalingWs))
}

func signalingWs(c *websocket.Conn) {
	localsToken := c.Locals("user").(*jwt.Token)
	claims := localsToken.Claims.(jwt.MapClaims)
	useridFloat, ok := claims["userid"].(float64)
	if !ok {
		c.Close()
		return
	}
	userid := int(useridFloat)
	connManager.mu.Lock()
	connManager.userid[userid] = &wrap{conn: c}
	connManager.mu.Unlock()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			connManager.mu.Lock()
			delete(connManager.userid, userid)
			connManager.mu.Unlock()
			c.Close()
			break
		}
		go msgHandler(msg)
	}
}

func msgHandler(msg []byte) {
	var message SignalMessage

	if err := json.Unmarshal(msg, &message); err != nil {
		log.Println("error! ", err)
		return
	}
	toInt, err := message.To.Int64()
	if err != nil {
		log.Println("couldn't convert field TO to int", err)
		return
	}
	connManager.mu.Lock()
	recConnection, exists := connManager.userid[int(toInt)]
	connManager.mu.Unlock()
	if !exists {
		log.Printf("couldn't find conn")
		return
	}
	recConnection.mu.Lock()
	if err := recConnection.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Println("failed to send message")
		connManager.mu.Lock()
		delete(connManager.userid, int(toInt))
		connManager.mu.Unlock()
		recConnection.conn.Close()
	}
	recConnection.mu.Unlock()
}
