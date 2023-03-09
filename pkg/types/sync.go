package types

import "github.com/gorilla/websocket"

type SyncClient struct {
	UserID   string
	ClientID string
	Conn     *websocket.Conn
}

type PubSubMessage struct {
	Type     string
	Action   string
	ClientID string
	Data     interface{}
}
