package models

import "fsm/pkg/utils"

// User is the model entity for the User schema.
type User struct {
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// PassWord holds the value of the "pass_word" field.
	PassWord string `json:"pass_word,omitempty"`
	// Salt holds the value of the "salt" field.
	Salt []byte `json:"salt,omitempty"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"user_name,omitempty"`
	// BucketName holds the value of the "bucket_name" field.
	BucketName string `json:"bucket_name,omitempty"`
	// CurrentStoreCap holds the value of the "current_store_cap" field.
	CurrentStoreCap int64 `json:"current_store_cap,omitempty"`
	// MaxStoreCap holds the value of the "max_store_cap" field.
	MaxStoreCap int64 `json:"max_store_cap,omitempty"`
}

// CheckPassword
func (u *User) CheckPassword(password string) bool {
	hash := utils.PasswordHash(password, u.Salt)
	return hash == u.PassWord
}

type UserFile struct {
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
