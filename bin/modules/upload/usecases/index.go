package usecases

import (
	"context"

	"github.com/Manusiabodoh4/bulk-generic/bin/entity"
	models "github.com/Manusiabodoh4/bulk-generic/bin/modules/upload/models/domain"
)

type UsecaseUploads interface {
	UploadFile(ctx context.Context, data *models.UploadRequest) entity.TemplateChannelResponse
}
