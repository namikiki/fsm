package repositories

import (
	"context"
	"fsm/models"
	"gorm.io/gorm"
)

type StorageRepository struct {
	db *gorm.DB
}

func NewStorageRepository(db *gorm.DB) *StorageRepository {
	return &StorageRepository{db: db}
}

func (sr *StorageRepository) GetFile(ctx context.Context, userID, fileID string) (*models.File, error) {
	var file models.File
	err := sr.db.Where("user_id = ? and id =?", userID, fileID).Find(&file).Error
	return &file, err
}

//func (sr *StorageRepository) FindAll(ctx context.Context, file *models.File) error {
//	return sr.db.Where(file).Find(file).Error
//}

func (sr *StorageRepository) CreateFile(ctx context.Context, file *models.File) error {
	return sr.db.Create(file).Error
}

func (sr *StorageRepository) UpdateFile(ctx context.Context, file *models.File) error {
	return sr.db.Save(file).Error
}

func (sr *StorageRepository) DeleteFile(ctx context.Context, userID, fileID string) error {
	var file models.File
	return sr.db.Where("user_id = ? and id =?", userID, fileID).Delete(&file).Error
}

//folder 文件夹

func (sr *StorageRepository) GetFolder(ctx context.Context, userID, folderID string) (*models.Folder, error) {
	var folder models.Folder
	err := sr.db.Where("user_id = ? and id =?", userID, folderID).Find(&folder).Error
	return &folder, err
}

func (sr *StorageRepository) ListFolder(ctx context.Context, userID string) ([]models.Folder, error) {
	var folder []models.Folder
	err := sr.db.Where("user_id = ?", userID, userID).Find(&folder).Error
	return folder, err
}

func (sr *StorageRepository) CreateFolder(ctx context.Context, folder *models.Folder) error {
	return sr.db.Create(folder).Error
}

func (sr *StorageRepository) UpdateFolder(ctx context.Context, folder *models.Folder) error {
	return sr.db.Save(folder).Error
}

func (sr *StorageRepository) DeleteFolder(ctx context.Context, userID, folderID string) error {
	var folder models.Folder
	return sr.db.Where("user_id = ? and id =?", userID, folderID).Delete(&folder).Error
}
