package types

import "time"

type File struct {
	SyncID     string    `json:"sync_id" form:"sync_id" validate:"required"`
	Name       string    `json:"file_name" form:"file_name" validate:"required"`
	ParentID   string    `json:"parent_id" form:"parent_id" validate:"required"`
	Level      uint64    `json:"level" form:"level" validate:"required"`
	CreateTime time.Time `json:"create_time" form:"create_time" validate:"required"`
	ModTIme    time.Time `json:"mod_time" json:"mod_time" validate:"required"`
	DeleteTime time.Time `json:"delete_time" form:"delete_time" validate:"required"`
}
