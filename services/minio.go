package services

import (
	"context"
	"fsm/models"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"time"
)

type MinioService struct {
	Min *minio.Client
}

// NewMinioService 创建一个新的 MinioService 实例
func NewMinioService(mio *minio.Client) *MinioService {
	return &MinioService{Min: mio}
}

// InitUserMinio 初始化用户的 MinIO 存储桶
func (m *MinioService) InitUserMinio(ctx context.Context, user *models.User) error {
	if err := m.Min.MakeBucket(ctx, user.ID, minio.MakeBucketOptions{ObjectLocking: false}); err != nil {
		log.Println("")
		return err
	}
	return nil
}

// FileCreate

// FileCreate 在 MinIO 中创建文件
func (m *MinioService) FileCreate(c *gin.Context, bucketName string, objectName string) (*minio.UploadInfo, error) {

	object, err := m.Min.PutObject(c, bucketName, objectName, c.Request.Body, c.Request.ContentLength,
		minio.PutObjectOptions{ContentType: "application/octet-stream"})

	if err != nil {
		//todo 日志
		log.Println("")
		return nil, err
	}

	return &object, nil
}

// FileDelete 删除 MinIO 中的文件
func (m *MinioService) FileDelete(ctx context.Context, bucketName, objectName string) error {
	return m.Min.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
}

// FileOpen 打开 MinIO 中的文件
func (m *MinioService) FileOpen(ctx context.Context, bucketName, fileID string) (
	io.ReadCloser, error) {
	return m.Min.GetObject(ctx, bucketName, fileID, minio.GetObjectOptions{})
}

// GetFileDownloadURL 获取文件下载链接
func (m *MinioService) GetFileDownloadURL(bucketName, fileName string,
	expires time.Duration) (string, error) {

	downloadURL, err := m.Min.PresignedGetObject(context.Background(), bucketName, fileName,
		expires, nil)
	return downloadURL.String(), err
}
