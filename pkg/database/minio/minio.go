package minio

import (
	"strings"

	"github.com/Nutchanon28/file-sharing-system/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient(config *config.Config) (*minio.Client, error) {
	minioClient, err := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretKey, ""),
		Secure: strings.ToLower(config.MinioUseSSL) == "true",
	})

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
