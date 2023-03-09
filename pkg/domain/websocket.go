package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketService interface {
	Run()
	Close() error
	Connect(c *gin.Context) error

	Info() map[string]map[string]*websocket.Conn
	BroadcastFileDown(uid string, clientUUID string)
}
