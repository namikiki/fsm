package services

import (
	"encoding/json"
	"fsm/models"
	"fsm/pkg/repositories"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"log"
	"time"
)

type Folder struct {
	ID         string `json:"id,omitempty"`          // 文件夹ID
	SyncID     string `json:"sync_id,omitempty"`     // 同步ID
	Path       string `json:"path,omitempty"`        // 文件夹路径
	Level      int    `json:"level,omitempty"`       // 文件夹级别
	Deleted    bool   `json:"deleted,omitempty"`     // 是否已删除
	CreateTime int64  `json:"create_time,omitempty"` // 创建时间
	ModTime    int64  `json:"mod_time,omitempty"`    // 修改时间
}

type FolderService struct {
	storageRepo  *repositories.StorageRepository
	MinioService *MinioService
	Redis        *redis.Client
}

// NewFolderService 创建一个新的 FolderService 实例
func NewFolderService(minioService *MinioService, redis *redis.Client, storageRepo *repositories.StorageRepository) *FolderService {
	return &FolderService{storageRepo, minioService, redis}
}

// ListFolder 列出用户的文件夹
func (f *FolderService) ListFolder(c *gin.Context, userID string) ([]models.Folder, error) {
	return f.storageRepo.ListFolder(c, userID)
}

// CreatFolder 创建新文件夹
func (f *FolderService) CreatFolder(c *gin.Context, folder Folder, userID, clientID string) (*models.Folder, error) {
	folder.ID = uuid.New().String()

	folderMeta := models.Folder{
		ID:         folder.ID,
		UserID:     userID,
		SyncID:     folder.SyncID,
		Path:       folder.Path,
		Level:      folder.Level,
		Deleted:    false,
		CreateTime: time.Now().Unix(),
		ModTime:    time.Now().Unix(),
	}

	if err := f.storageRepo.CreateFolder(c, &folderMeta); err != nil {
		log.Println(err)
		return nil, err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FOLDER, CREATE, folderMeta.SyncID, clientID, folderMeta))
	if err := f.Redis.Publish(c, userID, msg).Err(); err != nil {
		return nil, err
	}

	return &folderMeta, nil
}

// DeleteFolder 删除文件夹
func (f *FolderService) DeleteFolder(c *gin.Context, folderID, userID, clientID string) error {

	if err := f.storageRepo.DeleteFolder(c, userID, folderID); err != nil {
		log.Println(err)
		return err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FOLDER, DELETE, "", clientID, folderID))
	return f.Redis.Publish(c, userID, msg).Err()
}

// RenameFolder 重命名文件夹
func (f *FolderService) RenameFolder(c *gin.Context, folder Folder, userID, clientID string) error {

	getFolder, err := f.storageRepo.GetFolder(c, userID, folder.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	getFolder.Path = folder.Path
	if err := f.storageRepo.UpdateFolder(c, getFolder); err != nil {
		log.Println(err)
		return err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FOLDER, RENAME, folder.SyncID, clientID, getFolder))
	return f.Redis.Publish(c, userID, msg).Err()
}
