package services

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"log"
)

type SyncClient struct {
	UserID   string
	ClientID string
	Conn     *websocket.Conn
}

// WebSocketService 处理 WebSocket 连接和消息传递服务。
type WebSocketService struct {
	WebsocketConnChannel chan SyncClient
	Redis                *redis.Client
}

// NewWebSocketService 创建一个新的 WebSocketService 实例，并初始化 WebsocketConnChannel 和 Redis 客户端。
func NewWebSocketService(redis *redis.Client) *WebSocketService {
	return &WebSocketService{
		WebsocketConnChannel: make(chan SyncClient, 200),
		Redis:                redis,
	}
}

// HandleWebSocketConnections 处理 WebSocket 连接，在通道中接收到新的 SyncClient 时，启动一个新的 goroutine 来中继消息。
func (w *WebSocketService) HandleWebSocketConnections() {
	for {
		select {
		case wsc := <-w.WebsocketConnChannel:
			go w.RelayMessagesToClient(wsc)
		}
	}
}

// RelayMessagesToClient 从 Redis 订阅接收消息并中继给 WebSocket 客户端。
func (w *WebSocketService) RelayMessagesToClient(client SyncClient) {
	ctx := context.Background()
	subscribe := w.Redis.Subscribe(ctx, client.UserID)

	for {

		message, err := subscribe.ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
		}
		//var psm PubSubMessage
		//err = json.Unmarshal([]byte(message.Payload), &psm)

		if err := client.Conn.WriteMessage(websocket.BinaryMessage, []byte(message.Payload)); err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				client.Conn.Close()
				return
			}
		}

	}

}
