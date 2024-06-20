package models

// File 数据库文件结构
type File struct {
	ID          string `json:"id,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	SyncID      string `json:"sync_id,omitempty"`
	Name        string `json:"name,omitempty"`
	ParentDirID string `json:"parent_dir_id,omitempty"`
	Level       int    `json:"level,omitempty"`
	Hash        string `json:"hash,omitempty"`
	Size        int64  `json:"size,omitempty"`
	Deleted     bool   `json:"deleted,omitempty"`
	CreateTime  int64  `json:"create_time,omitempty"`
	ModTime     int64  `json:"mod_time,omitempty"`
}

// Folder 数据库文件夹结构
type Folder struct {
	ID         string `json:"id,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	SyncID     string `json:"sync_id,omitempty"`
	Path       string `json:"dir,omitempty"`
	Level      int    `json:"level,omitempty"`
	Deleted    bool   `json:"deleted,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	ModTime    int64  `json:"mod_time,omitempty"`
}
