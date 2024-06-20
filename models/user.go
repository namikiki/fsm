package models

import "fsm/pkg/utils"

// User  数据库用户结构
type User struct {
	ID       string `json:"id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Email    string `json:"email,omitempty"`
	PassWord string `json:"pass_word,omitempty"`
	Salt     []byte `json:"salt,omitempty"`
	Status   string `json:"status"`

	BucketName      string `json:"bucket_name,omitempty"`
	CurrentStoreCap int64  `json:"current_store_cap,omitempty"`
	MaxStoreCap     int64  `json:"max_store_cap,omitempty"`
}

func (u *User) CheckPassword(password string) bool {
	hash := utils.PasswordHash(password, u.Salt)
	return hash == u.PassWord
}
