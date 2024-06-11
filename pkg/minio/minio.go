package minio

import (
	"fsm/pkg/types"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

// NewMinioConnect 连接 MinIO 返回连接对象
func NewMinioConnect(config *types.Config) (*minio.Client, error) {
	minioClient, err := minio.New(config.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Minio.AccessKeyID, config.Minio.SecretAccessKey, ""),
		Secure: config.Minio.UseSSL,
	})

	if err != nil {
		log.Fatalln("minio 连接错误", err)
		return nil, err
	}
	return minioClient, nil
}
