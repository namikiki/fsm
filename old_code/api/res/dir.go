package res

import "time"

// Dir is the model entity for the Dir schema.
type Dir struct {
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	// SyncID holds the value of the "sync_id" field.
	SyncID string `json:"sync_id,omitempty"`
	// Dir holds the value of the "dir" field.
	Dir string `json:"dir,omitempty"`
	// Level holds the value of the "level" field.
	Level uint64 `json:"level,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// ModTime holds the value of the "mod_time" field.
	ModTime time.Time `json:"mod_time,omitempty"`
}
