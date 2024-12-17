package minio

import (
	"context"
	"mime/multipart"
	"regexp"

	"path/filepath"
	"time"

	"github.com/Nutchanon28/file-sharing-system/config"
	apiError "github.com/Nutchanon28/file-sharing-system/pkg/api_error"
	"github.com/Nutchanon28/file-sharing-system/pkg/constants"
	"github.com/Nutchanon28/file-sharing-system/pkg/database/minio/interfaces"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type MinioUsecase struct {
	config      *config.Config
	minioClient *minio.Client
	// logger      logger.Logger
}

func NewMinioUseCase(
	config *config.Config,
	minioClient *minio.Client,
	// logger logger.Logger,
) interfaces.MinioUseCase {

	return &MinioUsecase{
		config:      config,
		minioClient: minioClient,
		// logger:      logger,
	}
}

func (m *MinioUsecase) UploadFile(
	ctx context.Context,
	fileHeader *multipart.FileHeader,
	customName *string,
	filePolicy constants.FilePolicy,
	folder string,
) (*string, error) {
	passed, err := checkFilePolicy(fileHeader, filePolicy)

	if !passed || err != nil {
		// m.logger.Error(err)

		return nil, err
	}

	extension := filepath.Ext(fileHeader.Filename)

	var filename string
	if customName != nil {
		filename = folder + *customName + extension
	} else {
		filename = folder + uuid.New().String() + extension
	}

	file, err := fileHeader.Open()

	if err != nil {
		// m.logger.Error(err)

		return nil, apiError.NewBadRequestError("Damaged File")
	}

	defer file.Close()

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)

	defer cancel()

	_, err = m.minioClient.PutObject(
		ctx,
		m.config.MinioBucketName,
		filename,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{},
	)

	if err != nil {
		// m.logger.Error(err)

		return nil, apiError.NewServiceUnavaliableError("Minio service unavailable")
	}

	return &filename, nil
}

func (m *MinioUsecase) Move(
	ctx context.Context,
	dst string,
	src string,
) error {
	err := m.copy(ctx, dst, src)

	if err != nil {
		return err
	}

	return m.remove(ctx, src)
}

func (m *MinioUsecase) GetPresignedUrl(ctx context.Context, objectName string) (*string, error) {
	duration := time.Duration(m.config.MinioPresignedDuration) * time.Hour

	presignedUrl, err := m.minioClient.PresignedGetObject(
		ctx,
		m.config.MinioBucketName,
		objectName,
		duration,
		nil,
	)

	if err != nil {
		return nil, err
	}

	url := presignedUrl.String()

	return &url, nil
}

func checkFilePolicy(fileHeader *multipart.FileHeader, filePolicy constants.FilePolicy) (bool, error) {
	mimeType := fileHeader.Header.Get("Content-Type")
	isExceed := fileHeader.Size > int64(filePolicy.MaxSize)

	matched, err := regexp.MatchString(filePolicy.FileType, mimeType)
	if !matched || err != nil {
		return false, apiError.NewBadRequestError("File mimetype mismatch")
	}

	if isExceed {
		return false, apiError.NewBadRequestError("File size exceeded")
	}

	return true, nil
}

func (m *MinioUsecase) copy(
	ctx context.Context,
	dst string,
	src string,
) error {
	srcOpts := minio.CopySrcOptions{
		Bucket: m.config.MinioBucketName,
		Object: src,
	}

	dstOpts := minio.CopyDestOptions{
		Bucket: m.config.MinioBucketName,
		Object: dst,
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	_, err := m.minioClient.CopyObject(ctx, dstOpts, srcOpts)

	if err != nil {
		// m.logger.Error(err)

		return apiError.NewServiceUnavaliableError("Minio service unavailable")
	}

	return nil
}

func (m *MinioUsecase) remove(
	ctx context.Context,
	filename string,
) error {
	opts := minio.RemoveObjectOptions{}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	err := m.minioClient.RemoveObject(
		ctx,
		m.config.MinioBucketName,
		filename,
		opts,
	)

	if err != nil {
		// m.logger.Error(err)

		return apiError.NewServiceUnavaliableError("Minio service unavailable")
	}

	return nil
}
