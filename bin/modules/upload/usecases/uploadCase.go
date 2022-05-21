package usecases

import (
	"context"
	"errors"

	"github.com/Manusiabodoh4/bulk-generic/bin/entity"
	models "github.com/Manusiabodoh4/bulk-generic/bin/modules/upload/models/domain"
	"github.com/Manusiabodoh4/bulk-generic/bin/modules/upload/repositories"
	"github.com/Manusiabodoh4/bulk-generic/bin/pkg/databases"
	"github.com/xuri/excelize/v2"
)

type UsecaseUploadsImpl struct {
	uploadRepo     repositories.UploadBulkConnectionRepo
	bulkConnection *databases.BulkConnectionPkg
}

func NewUsecaseUploads(uploadRepo repositories.UploadBulkConnectionRepo, bulkConnection databases.BulkConnectionPkg) UsecaseUploads {
	return &UsecaseUploadsImpl{
		uploadRepo:     uploadRepo,
		bulkConnection: &bulkConnection,
	}
}

func (h *UsecaseUploadsImpl) UploadFile(ctx context.Context, data *models.UploadRequest) entity.TemplateChannelResponse {
	var (
		result     entity.TemplateChannelResponse
		lenData    int
		lenField   int
		payload    []map[string]interface{}
		tmpPayload map[string]interface{}
	)

	f, err := excelize.OpenReader(data.File)
	if err != nil {
		result.Error = err
		return result
	}
	defer f.Close()

	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		result.Error = err
		return result
	}

	lenData = len(rows) - 1
	lenField = len(rows[0])
	if lenData <= 0 || lenField <= 0 {
		result.Error = errors.New("tidak terdapat data didalam file excel")
		return result
	}

	for i := 1; i <= lenData; i++ {
		tmpPayload = make(map[string]interface{})
		for j := 0; j < lenField; j++ {
			tmpPayload[rows[0][j]] = rows[i][j]
		}
		payload = append(payload, tmpPayload)
	}

	isFound, tx := h.bulkConnection.GetBulkConnection(data.ID)
	if !isFound {
		result.Error = errors.New("connection dengan id " + data.ID + " tidak ditemukan")
		return result
	}
	queryRes := <-h.uploadRepo.InsertMany(ctx, tx.Table(data.TableName), &payload)
	if queryRes.Error != nil {
		result.Error = queryRes.Error
		return result
	}

	result.Data = queryRes.Data
	return result
}
