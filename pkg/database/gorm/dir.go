package gorm

import (
	"context"

	"fsm/pkg/domain"
	"fsm/pkg/ent"

	"gorm.io/gorm"
)

type DirRepository struct {
	Conn *gorm.DB
}

func NewDirRepository(conn *gorm.DB) domain.DirRepository {
	return &DirRepository{Conn: conn}
}

func (d *DirRepository) Create(ctx context.Context, dir ent.Dir) error {
	d.Conn.Create(&dir)
	return nil
}

func (d *DirRepository) Delete(ctx context.Context, dir ent.Dir) error {
	d.Conn.Where("id = ? and sync_id = ?", dir.ID, dir.SyncID).Delete(&dir)
	return nil
}

func (d *DirRepository) Rename(ctx context.Context, dir ent.Dir) error {
	d.Conn.Save(&dir)
	return nil
}

func (d *DirRepository) ReadDir(ctx context.Context, dir ent.Dir) ([]ent.Dir, error) {
	var dirs []ent.Dir
	d.Conn.Where("user_id=? and sync_id=? and level=? and parent_dir=?", dir.UserID,
		dir.SyncID, dir.Level+1, dir.Dir).Find(&dirs)
	return dirs, nil
}

func (d *DirRepository) WalkDirByPath(ctx context.Context, dir ent.Dir) ([]ent.Dir, error) {
	var dirs []ent.Dir
	d.Conn.Where("user_id = ? and sync_id = ? and level > ? and parent_dir like ?",
		dir.UserID, dir.SyncID, dir.Level, dir.Dir+"%").Find(&dirs)
	return dirs, nil
}

func (d *DirRepository) WalkDirBySyncID(ctx context.Context, userID, syncID string) ([]ent.Dir, error) {
	var dirs []ent.Dir
	d.Conn.Where("user_id = ? and sync_id =?", userID, syncID).Find(&dirs)
	return dirs, nil
}
