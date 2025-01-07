package signaling

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

type SignalMessage struct {
	Type string          `json:"type"`
	From json.Number     `json:"from"`
	To   json.Number     `json:"to"`
	Data json.RawMessage `json:"data"`
}

type wrap struct {
	conn  *websocket.Conn
	mu    sync.Mutex
	msgCh chan []byte
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
	connManager.userid[userid] = &wrap{
		conn:  c,
		msgCh: make(chan []byte, 50),
	}
	msgCh := connManager.userid[userid].msgCh
	connManager.mu.Unlock()
	go msgHandler(msgCh, userid)

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			connManager.mu.Lock()
			delete(connManager.userid, userid)
			connManager.mu.Unlock()
			c.Close()
			break
		}
		msgCh <- msg
	}
}

func msgHandler(msgCh chan []byte, userid int) {
	var message SignalMessage
	for msg := range msgCh {
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("error! ", err)
			return
		}
		//	log.Println(message.To) it sometimes skips ice i think, idk =)
		if message.To == "" {
			continue
		}
		toInt, err := message.To.Int64()
		if err != nil {
			log.Println("couldn't convert field TO to int", err)
			continue
		}
		if message.From == "" {
			message.From = json.Number(strconv.Itoa(userid))
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
			close(recConnection.msgCh)
		}
		recConnection.mu.Unlock()
	}
}
