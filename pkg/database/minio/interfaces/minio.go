package interfaces

import (
	"context"
	"mime/multipart"

	"github.com/Nutchanon28/file-sharing-system/pkg/constants"
)

type MinioUseCase interface {
	UploadFile(
		ctx context.Context,
		fileHeader *multipart.FileHeader,
		customName *string,
		filePolicy constants.FilePolicy,
		folder string,
	) (*string, error)
	GetPresignedUrl(ctx context.Context, objectName string) (*string, error)
	Move(ctx context.Context, dst string, src string) error
}
