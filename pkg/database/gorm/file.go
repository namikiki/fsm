package gorm

import (
	"context"

	"fsm/api/res"
	"fsm/pkg/domain"
	"fsm/pkg/ent"

	"gorm.io/gorm"
)

type FileRepository struct {
	Conn *gorm.DB
}

func NewFileRepository(conn *gorm.DB) domain.FileRepository {
	return &FileRepository{Conn: conn}
}

func (fr *FileRepository) GetMetadataByID(ctx context.Context, userID, syncID, fileID string) (ent.File, error) {
	var f ent.File
	fr.Conn.Where("user_id = ? and sync_id = ? and id =?", userID, syncID, fileID).Find(&f)
	return f, nil
}

func (fr *FileRepository) Create(ctx context.Context, f *ent.File) error {
	fr.Conn.Create(f)
	return nil
}

func (fr *FileRepository) Delete(ctx context.Context, f ent.File) error {
	fr.Conn.Delete(&f)
	return nil
}

func (fr *FileRepository) Update(ctx context.Context, f ent.File) error {
	fr.Conn.Save(&f)
	return nil
}

func (fr *FileRepository) GetAllBySyncID(ctx context.Context, userID, syncID string) ([]res.File, error) {
	var files []res.File
	fr.Conn.Where("user_id =? and sync_id = ?", userID, syncID).Find(&files)
	return files, nil
}
