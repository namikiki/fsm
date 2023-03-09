package minio

import (
	"context"
	"io"
	"time"

	"fsm/pkg/domain"

	"github.com/minio/minio-go/v7"
)

type minioFileStorageService struct {
	Min *minio.Client
}

func NewMinioFileStorageService(mio *minio.Client) domain.FileStorageService {
	return &minioFileStorageService{Min: mio}
}

func (m *minioFileStorageService) Create(ctx context.Context, uid, parentId, name string,
	read io.Reader, length int64) error {

	return nil
}

func (m *minioFileStorageService) Delete(ctx context.Context, bucketName, fileName string) error {
	return m.Min.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
}

func (m *minioFileStorageService) Open(ctx context.Context, bucketName, fileID string) (
	io.ReadCloser, error) {
	return m.Min.GetObject(ctx, bucketName, fileID, minio.GetObjectOptions{})
}

func (m *minioFileStorageService) GetDownloadURL(bucketName, fileName string,
	expires time.Duration) (string, error) {

	downloadURL, err := m.Min.PresignedGetObject(context.Background(), bucketName, fileName,
		expires, nil)
	return downloadURL.String(), err
}

func (m *minioFileStorageService) Update(ctx context.Context, BucketName, ID string,
	Read io.Reader, length int64) error {
	//TODO implement me
	panic("implement me")
}
