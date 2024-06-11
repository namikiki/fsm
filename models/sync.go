package models

type PubSubMessage struct {
	Type     string      `json:"type"`
	Action   string      `json:"action"`
	SyncID   string      `json:"sync_id"`
	ClientID string      `json:"client_id"`
	Data     interface{} `json:"data"`
}

func NewPubSubMessage(types, action, syncID, clientID string, data any) PubSubMessage {
	return PubSubMessage{types, action, syncID, clientID, data}
}
