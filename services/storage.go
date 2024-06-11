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

type StorageService struct {
	storageRepo  *repositories.StorageRepository
	MinioService *MinioService
	Redis        *redis.Client
}

type File struct {
	ID          string `json:"id"`
	SyncID      string `json:"sync_id,omitempty"`
	Name        string `json:"name,omitempty"`
	ParentDirID string `json:"parent_dir_id,omitempty"`
	Level       int    `json:"level,omitempty"`
	CreateTime  int64  `json:"create_time,omitempty"`
	ModTime     int64  `json:"mod_time,omitempty"`
}

type Folder struct {
	ID         string `json:"id,omitempty"`
	SyncID     string `json:"sync_id,omitempty"`
	Path       string `json:"path,omitempty"`
	Level      int    `json:"level,omitempty"`
	Deleted    bool   `json:"deleted,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	ModTime    int64  `json:"mod_time,omitempty"`
}

func NewStorageService(minioService *MinioService, redis *redis.Client, storageRepo *repositories.StorageRepository) *StorageService {
	return &StorageService{storageRepo, minioService, redis}
}

func (s *StorageService) CreateFile(c *gin.Context, file File, userID, clientID string) (*models.File, error) {
	file.ID = uuid.New().String()
	info, err := s.MinioService.FileCreate(c, userID, file.ID)
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

	if err := s.storageRepo.CreateFile(c, &fileMeta); err != nil {
		log.Println(err)
		return nil, err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FILE, CREATE, fileMeta.SyncID, clientID, fileMeta))
	if err := s.Redis.Publish(c, userID, msg).Err(); err != nil {
		return nil, err
	}

	return &fileMeta, nil
}

func (s *StorageService) DeleteFile(c *gin.Context, fileID, userID, clientID string) error {

	if err := s.storageRepo.DeleteFile(c, userID, fileID); err != nil {
		log.Println(err)
		return err
	}

	if err := s.MinioService.FileDelete(c, userID, fileID); err != nil {
		log.Println(err)
		return err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FILE, DELETE, "", clientID, fileID))
	return s.Redis.Publish(c, userID, msg).Err()
}

func (s *StorageService) RenameFile(c *gin.Context, file File, userID, clientID string) error {
	getFile, err := s.storageRepo.GetFile(c, userID, file.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	getFile.Name = file.Name
	if err := s.storageRepo.UpdateFile(c, getFile); err != nil {
		log.Println(err)
		return err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FILE, RENAME, file.SyncID, clientID, getFile))
	return s.Redis.Publish(c, userID, msg).Err()
}

func (s *StorageService) UpdateFile(c *gin.Context, file File, userID, clientID string) error {
	getFile, err := s.storageRepo.GetFile(c, userID, file.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	info, err := s.MinioService.FileCreate(c, userID, file.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	getFile.Size = info.Size
	getFile.Hash = info.ETag
	getFile.ModTime = info.LastModified.Unix()

	if err := s.storageRepo.UpdateFile(c, getFile); err != nil {
		log.Println(err)
		return err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FILE, UPDATE, file.SyncID, clientID, getFile))
	return s.Redis.Publish(c, userID, msg).Err()
}

func (s *StorageService) GetFileContent(c *gin.Context, fileID, userID string) (io.ReadCloser, error) {
	getFile, err := s.storageRepo.GetFile(c, userID, fileID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return s.MinioService.FileOpen(c, userID, getFile.ID)
}

func (s *StorageService) GetFileMeta(c *gin.Context, fileID, userID string) (*models.File, error) {
	return s.storageRepo.GetFile(c, userID, fileID)
}

//folder

func (s *StorageService) ListFolder(c *gin.Context, userID string) ([]models.Folder, error) {
	return s.storageRepo.ListFolder(c, userID)
}

func (s *StorageService) CreatFolder(c *gin.Context, folder Folder, userID, clientID string) (*models.Folder, error) {
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

	if err := s.storageRepo.CreateFolder(c, &folderMeta); err != nil {
		log.Println(err)
		return nil, err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FOLDER, CREATE, folderMeta.SyncID, clientID, folderMeta))
	if err := s.Redis.Publish(c, userID, msg).Err(); err != nil {
		return nil, err
	}

	return &folderMeta, nil
}

func (s *StorageService) DeleteFolder(c *gin.Context, folderID, userID, clientID string) error {

	if err := s.storageRepo.DeleteFolder(c, userID, folderID); err != nil {
		log.Println(err)
		return err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FOLDER, DELETE, "", clientID, folderID))
	return s.Redis.Publish(c, userID, msg).Err()
}

func (s *StorageService) RenameFolder(c *gin.Context, folder Folder, userID, clientID string) error {

	getFolder, err := s.storageRepo.GetFolder(c, userID, folder.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	getFolder.Path = folder.Path
	if err := s.storageRepo.UpdateFolder(c, getFolder); err != nil {
		log.Println(err)
		return err
	}

	msg, _ := json.Marshal(models.NewPubSubMessage(FOLDER, RENAME, folder.SyncID, clientID, getFolder))
	return s.Redis.Publish(c, userID, msg).Err()
}
