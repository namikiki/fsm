package repositories

import (
	"context"
	"fsm/models"
	"gorm.io/gorm"
)

type StorageRepository struct {
	db *gorm.DB
}

// NewStorageRepository 使用提供的 gorm.DB 实例创建一个新的 StorageRepository 实例。
func NewStorageRepository(db *gorm.DB) *StorageRepository {
	return &StorageRepository{db: db}
}

// GetFile 根据文件ID检索特定用户的文件元信息
func (sr *StorageRepository) GetFile(ctx context.Context, userID, fileID string) (*models.File, error) {
	var file models.File
	err := sr.db.Where("user_id = ? and id =?", userID, fileID).Find(&file).Error
	return &file, err
}

//func (sr *StorageRepository) FindAll(ctx context.Context, file *models.File) error {
//	return sr.db.Where(file).Find(file).Error
//}

// CreateFile 在数据库中创建一个新的文件记录。
func (sr *StorageRepository) CreateFile(ctx context.Context, file *models.File) error {
	return sr.db.Create(file).Error
}

// UpdateFile 更新数据库中现有的文件记录。
func (sr *StorageRepository) UpdateFile(ctx context.Context, file *models.File) error {
	return sr.db.Save(file).Error
}

// DeleteFile 根据文件ID删除特定用户的文件。
func (sr *StorageRepository) DeleteFile(ctx context.Context, userID, fileID string) error {
	var file models.File
	return sr.db.Where("user_id = ? and id =?", userID, fileID).Delete(&file).Error
}

//folder 文件夹

// GetFolder 根据文件夹ID检索特定用户的文件夹。
func (sr *StorageRepository) GetFolder(ctx context.Context, userID, folderID string) (*models.Folder, error) {
	var folder models.Folder
	err := sr.db.Where("user_id = ? and id =?", userID, folderID).Find(&folder).Error
	return &folder, err
}

// ListFolder 列出特定用户的所有文件夹。
func (sr *StorageRepository) ListFolder(ctx context.Context, userID string) ([]models.Folder, error) {
	var folder []models.Folder
	err := sr.db.Where("user_id = ?", userID, userID).Find(&folder).Error
	return folder, err
}

// CreateFolder 在数据库中创建一个新的文件夹记录。
func (sr *StorageRepository) CreateFolder(ctx context.Context, folder *models.Folder) error {
	return sr.db.Create(folder).Error
}

// UpdateFolder 更新数据库中现有的文件夹记录。
func (sr *StorageRepository) UpdateFolder(ctx context.Context, folder *models.Folder) error {
	return sr.db.Save(folder).Error
}

// DeleteFolder 根据文件夹ID删除特定用户的文件夹。
func (sr *StorageRepository) DeleteFolder(ctx context.Context, userID, folderID string) error {
	var folder models.Folder
	return sr.db.Where("user_id = ? and id =?", userID, folderID).Delete(&folder).Error
}
