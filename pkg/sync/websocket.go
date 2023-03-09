package sync

import (
	"context"
	"log"

	"fsm/pkg/types"

	"github.com/gorilla/websocket"
)

func (s *Syncer) WebSocketLoop() {
	for {
		select {
		case w := <-s.WebsocketConnChannel:
			go s.SubAndSendClient(w)
		}
	}
}

func (s *Syncer) SubAndSendClient(c types.SyncClient) {
	ctx := context.Background()
	subscribe := s.Redis.Subscribe(ctx, c.UserID)

	for {
		message, err := subscribe.ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
		}
		//var psm PubSubMessage
		//err = json.Unmarshal([]byte(message.Payload), &psm)

		if err := c.Conn.WriteMessage(websocket.BinaryMessage, []byte(message.Payload)); err != nil {
			log.Println(err)
			continue
		}
	}

}
