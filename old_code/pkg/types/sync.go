package types

import "github.com/gorilla/websocket"

type SyncClient struct {
	UserID   string
	ClientID string
	Conn     *websocket.Conn
}

type PubSubMessage struct {
	Type     string `json:"type"`
	Action   string `json:"action"`
	SyncID   string `json:"sync_id"`
	ClientID string `json:"client_id"`
	Data     []byte `json:"data"`
}
