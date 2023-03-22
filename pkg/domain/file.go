package domain

import (
	"context"
	"io"
	"time"

	"fsm/pkg/ent"
)

type FileRepository interface {
	Create(ctx context.Context, f *ent.File) error
	Delete(ctx context.Context, f ent.File) error
	GetMetadataByID(ctx context.Context, userID, fileID string) (ent.File, error)
	Update(ctx context.Context, f ent.File) error
	Rename(ctx context.Context, f ent.File) error
	GetAllBySyncID(ctx context.Context, userID, syncID string) ([]ent.File, error)
}

type FileStorageService interface {
	Create(ctx context.Context, uid, parentId, name string, Read io.Reader, length int64) error
	Delete(ctx context.Context, bucketName, fileName string) error
	Open(ctx context.Context, bucketName, fileID string) (io.ReadCloser, error)
	Update(ctx context.Context, BucketName, ID string, Read io.Reader, length int64) error
	GetDownloadURL(bucketName, fileName string, expires time.Duration) (string, error)
}
