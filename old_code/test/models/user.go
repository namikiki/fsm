package models

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
