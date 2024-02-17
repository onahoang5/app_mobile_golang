package uploadprovider

import (
	"Food-delivery/common"
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(context context.Context, data []byte, dst string) (*common.Image, error)
}
