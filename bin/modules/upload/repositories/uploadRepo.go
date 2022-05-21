package repositories

import (
	"context"

	"github.com/Manusiabodoh4/bulk-generic/bin/entity"
	"gorm.io/gorm"
)

type UploadBulkConnectionRepoImpl struct{}

func NewUploadBulkConnectionRepo() UploadBulkConnectionRepo {
	return &UploadBulkConnectionRepoImpl{}
}

func (h *UploadBulkConnectionRepoImpl) InsertMany(ctx context.Context, tx *gorm.DB, data *[]map[string]interface{}) <-chan entity.TemplateChannelResponse {
	output := make(chan entity.TemplateChannelResponse)

	go func() {
		defer close(output)
		res := tx.Create(data)
		if res.Error != nil {
			output <- entity.TemplateChannelResponse{
				Error: res.Error,
			}
		}
		output <- entity.TemplateChannelResponse{
			Data: data,
		}
	}()

	return output
}
