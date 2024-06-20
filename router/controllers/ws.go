package controllers

import (
	"fsm/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebsocketController struct {
	webSocketService *services.WebSocketService
}

// WebSocket 升级器，设置读写缓冲区大小
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// NewWebsocketController 创建一个新的 WebsocketController 实例
func NewWebsocketController(webSocketService *services.WebSocketService) *WebsocketController {
	return &WebsocketController{webSocketService}
}

// WebsocketConn 处理 WebSocket 连接
func (w *WebsocketController) WebsocketConn(c *gin.Context) {
	// 升级 HTTP 连接为 WebSocket 连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// 如果升级失败，返回错误信息
		ErrorResponse(c, 1, "WebSocket 连接失败", err)
	}

	// 创建一个新的 WebSocket 客户端
	wsClient := services.SyncClient{
		UserID:   c.GetHeader("userID"),
		ClientID: c.GetHeader("clientID"),
		Conn:     conn,
	}

	// 将 WebSocket 客户端放入 WebSocket 连接通道中
	w.webSocketService.WebsocketConnChannel <- wsClient
}
