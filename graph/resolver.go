package graph

import (
	"github.com/Nutchanon28/file-sharing-system/graph/model"
	minioInterfaces "github.com/Nutchanon28/file-sharing-system/pkg/database/minio/interfaces"
	"github.com/minio/minio-go/v7"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	SharedFilesList []*model.SharedFile
	MinioUseCase    minioInterfaces.MinioUseCase
	MinioClient     *minio.Client
}
