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

func NewMinioService(mio *minio.Client) *MinioService {
	return &MinioService{Min: mio}
}

func (m *MinioService) InitUserMinio(ctx context.Context, user *models.User) error {
	if err := m.Min.MakeBucket(ctx, user.ID, minio.MakeBucketOptions{ObjectLocking: false}); err != nil {
		log.Println("")
		return err
	}
	return nil
}

// FileCreate
// todo
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

func (m *MinioService) FileDelete(ctx context.Context, bucketName, objectName string) error {
	return m.Min.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
}

func (m *MinioService) FileOpen(ctx context.Context, bucketName, fileID string) (
	io.ReadCloser, error) {
	return m.Min.GetObject(ctx, bucketName, fileID, minio.GetObjectOptions{})
}

func (m *MinioService) GetFileDownloadURL(bucketName, fileName string,
	expires time.Duration) (string, error) {

	downloadURL, err := m.Min.PresignedGetObject(context.Background(), bucketName, fileName,
		expires, nil)
	return downloadURL.String(), err
}
