package services

import (
	"encoding/json"
	"fsm/models"
	"fsm/pkg/repositories"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"io"
	"log"
	"time"
)

// 利用 minio 服务和redis对象 为 控制层 fs 提供服务
const (
	FILE   = "file"
	FOLDER = "folder"
	CREATE = "create"
	DELETE = "delete"
	UPDATE = "delete"
	RENAME = "rename"
)

type FileService struct {
	storageRepo  *repositories.StorageRepository
	MinioService *MinioService
	Redis        *redis.Client
}

type File struct {
	ID          string `json:"id"`                      // 文件ID
	SyncID      string `json:"sync_id,omitempty"`       // 同步ID
	Name        string `json:"name,omitempty"`          // 文件名称
	ParentDirID string `json:"parent_dir_id,omitempty"` // 父目录ID
	Level       int    `json:"level,omitempty"`         // 文件级别
	CreateTime  int64  `json:"create_time,omitempty"`   // 创建时间
	ModTime     int64  `json:"mod_time,omitempty"`      // 修改时间
}

// NewFileService 创建一个新的 StorageService 实例
func NewFileService(minioService *MinioService, redis *redis.Client, storageRepo *repositories.StorageRepository) *FileService {
	return &FileService{storageRepo, minioService, redis}
}

// CreateFile 创建新文件
func (f *FileService) CreateFile(c *gin.Context, file File, userID, clientID string) (*models.File, error) {
	file.ID = uuid.New().String()
	info, err := f.MinioService.FileCreate(c, userID, file.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fileMeta := models.File{
		ID:          file.ID,
		UserID:      userID,
		SyncID:      file.SyncID,
		Name:        file.Name,
		ParentDirID: file.ParentDirID,
		Level:       file.Level,
		Hash:        info.ETag,
		Size:        info.Size,
		Deleted:     false,
		CreateTime:  time.Now().Unix(),
		ModTime:     time.Now().Unix(),
	}

	if err := f.storageRepo.CreateFile(c, &fileMeta); err != nil {
		log.Println(err)
		return nil, err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FILE, CREATE, fileMeta.SyncID, clientID, fileMeta))
	if err := f.Redis.Publish(c, userID, msg).Err(); err != nil {
		return nil, err
	}

	return &fileMeta, nil
}

// DeleteFile 删除文件
func (f *FileService) DeleteFile(c *gin.Context, fileID, userID, clientID string) error {

	if err := f.storageRepo.DeleteFile(c, userID, fileID); err != nil {
		log.Println(err)
		return err
	}

	if err := f.MinioService.FileDelete(c, userID, fileID); err != nil {
		log.Println(err)
		return err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FILE, DELETE, "", clientID, fileID))
	return f.Redis.Publish(c, userID, msg).Err()
}

// RenameFile 重命名文件
func (f *FileService) RenameFile(c *gin.Context, file File, userID, clientID string) error {
	getFile, err := f.storageRepo.GetFile(c, userID, file.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	getFile.Name = file.Name
	if err := f.storageRepo.UpdateFile(c, getFile); err != nil {
		log.Println(err)
		return err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FILE, RENAME, file.SyncID, clientID, getFile))
	return f.Redis.Publish(c, userID, msg).Err()
}

// UpdateFile 更新文件
func (f *FileService) UpdateFile(c *gin.Context, file File, userID, clientID string) error {
	getFile, err := f.storageRepo.GetFile(c, userID, file.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	info, err := f.MinioService.FileCreate(c, userID, file.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	getFile.Size = info.Size
	getFile.Hash = info.ETag
	getFile.ModTime = info.LastModified.Unix()

	if err := f.storageRepo.UpdateFile(c, getFile); err != nil {
		log.Println(err)
		return err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FILE, UPDATE, file.SyncID, clientID, getFile))
	return f.Redis.Publish(c, userID, msg).Err()
}

// GetFileContent 获取文件内容
func (f *FileService) GetFileContent(c *gin.Context, fileID, userID string) (io.ReadCloser, error) {
	getFile, err := f.storageRepo.GetFile(c, userID, fileID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return f.MinioService.FileOpen(c, userID, getFile.ID)
}

// GetFileMeta 获取文件元数据
func (f *FileService) GetFileMeta(c *gin.Context, fileID, userID string) (*models.File, error) {
	return f.storageRepo.GetFile(c, userID, fileID)
}
