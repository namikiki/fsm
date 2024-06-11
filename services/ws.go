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

type WebSocketService struct {
	WebsocketConnChannel chan SyncClient
	Redis                *redis.Client
}

func NewWebSocketService(redis *redis.Client) *WebSocketService {
	return &WebSocketService{
		WebsocketConnChannel: make(chan SyncClient, 200),
		Redis:                redis,
	}
}

func (w *WebSocketService) HandleWebSocketConnections() {
	for {
		select {
		case wsc := <-w.WebsocketConnChannel:
			go w.RelayMessagesToClient(wsc)
		}
	}
}

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
