package minio

import (
	"log"

	"fsm/pkg/types"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioConnect(config *types.Config) (*minio.Client, error) {

	//endpoint  = "play.min.io"
	//accessKey = "Q3AM3UQ867SPQQA43P2F"
	//secretKey = "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	//useSSL    = true
	//bucket    = "test-bucket"

	endpoint := config.Minio.Endpoint
	accessKeyID := config.Minio.AccessKeyID
	secretAccessKey := config.Minio.SecretAccessKey
	useSSL := config.Minio.UseSSL

	//endpoint := "play.min.io"
	//accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	//secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	//useSSL := true

	//endpoint := "127.0.0.1:9000"
	//accessKeyID := "minioadmin"
	//secretAccessKey := "minioadmin"
	//useSSL := false

	// Initialize minio client object.
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
