package minio

import (
	"log"

	"fsm/pkg/types"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioConnect(config *types.Config) (*minio.Client, error) {
	endpoint := config.Minio.Endpoint
	accessKeyID := config.Minio.AccessKeyID
	secretAccessKey := config.Minio.SecretAccessKey
	useSSL := config.Minio.UseSSL

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln("minio 连接错误", err)
		return nil, err
	}

	return minioClient, nil
}
