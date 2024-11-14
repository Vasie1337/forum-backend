package handlers

import (
	"encoding/json"
	"log"
	"server/internal/message"
	"server/internal/services"
	"server/internal/websocket"

	"github.com/gin-gonic/gin"
)

type LoaderHandler struct {
	LoaderService *services.LoaderService
}

func NewLoaderHandler(loaderService *services.LoaderService) *LoaderHandler {
	return &LoaderHandler{
		LoaderService: loaderService,
	}
}

func (h *LoaderHandler) SocketHandler(c *gin.Context) {
	conn, err := websocket.Upgrader.Upgrade(c.Writer, c.Request, nil)
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

		var msg message.Msg
		err = json.Unmarshal(p, &msg)
		if err != nil {
			log.Println(err)
			return
		}

		switch msg.Type {
		case message.ExchaneKeys:
			// exchange keys
		case message.UserLogin:
			// user login
		case message.DownloadCheat:
			// download cheat
		case message.Heartbeat:
			// heartbeat
		}
	}
}
