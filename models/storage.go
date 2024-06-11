package models

// File is the model entity for the File schema.
type File struct {
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// SyncID holds the value of the "sync_id" field.
	SyncID string `json:"sync_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ParentDirID holds the value of the "parent_dir_id" field.
	ParentDirID string `json:"parent_dir_id,omitempty"`
	// Level holds the value of the "level" field.
	Level int `json:"level,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"hash,omitempty"`
	// Size holds the value of the "size" field.
	Size int64 `json:"size,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime int64 `json:"create_time,omitempty"`
	// ModTime holds the value of the "mod_time" field.
	ModTime int64 `json:"mod_time,omitempty"`
}

// Folder is the model entity for the Folder schema.
type Folder struct {
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// SyncID holds the value of the "sync_id" field.
	SyncID string `json:"sync_id,omitempty"`
	// Dir holds the value of the "dir" field.
	Path string `json:"dir,omitempty"`
	// Level holds the value of the "level" field.
	Level int `json:"level,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime int64 `json:"create_time,omitempty"`
	// ModTime holds the value of the "mod_time" field.
	ModTime int64 `json:"mod_time,omitempty"`
}
