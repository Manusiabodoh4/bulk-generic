package repositories

import (
	"context"

	"github.com/Manusiabodoh4/bulk-generic/bin/entity"
	"gorm.io/gorm"
)

type UploadBulkConnectionRepo interface {
	InsertMany(ctx context.Context, tx *gorm.DB, data *[]map[string]interface{}) <-chan entity.TemplateChannelResponse
}
