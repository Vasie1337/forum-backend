package websocket

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type MessageType int

const (
	ExchaneKeys MessageType = iota
	UserLogin
	DownloadCheat
	Heartbeat
)

type Message struct {
	Type MessageType `json:"type"`
	Hash string      `json:"Hash"`
	Data interface{} `json:"data"`
}

func LoaderHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var msg Message
		err = json.Unmarshal(p, &msg)
		if err != nil {
			log.Println(err)
			return
		}

		switch msg.Type {
		case ExchaneKeys:
			// exchange keys
		case UserLogin:
			// user login
		case DownloadCheat:
			// download cheat
		case Heartbeat:
			// heartbeat
		}
	}
}
