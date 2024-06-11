package controllers

import (
	"fsm/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketController struct {
	webSocketService *services.WebSocketService
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewWebsocketController(webSocketService *services.WebSocketService) *WebsocketController {
	return &WebsocketController{webSocketService}
}

func (w *WebsocketController) WebsocketConn(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithStatusJSON(1003, NewErrorApiResult(501, err.Error()))
		return
	}

	wsClient := services.SyncClient{
		UserID:   c.GetHeader("userID"),
		ClientID: c.GetHeader("clientID"),
		Conn:     conn,
	}

	w.webSocketService.WebsocketConnChannel <- wsClient
}
