package handlers

import (
	"context"

	models "github.com/Manusiabodoh4/bulk-generic/bin/modules/upload/models/domain"
	"github.com/Manusiabodoh4/bulk-generic/bin/modules/upload/repositories"
	"github.com/Manusiabodoh4/bulk-generic/bin/modules/upload/usecases"
	"github.com/Manusiabodoh4/bulk-generic/bin/pkg/databases"
	"github.com/Manusiabodoh4/bulk-generic/bin/utils"
	"github.com/labstack/echo/v4"
)

type HttHandlersUpload struct {
	Logger   utils.ToolsLogger
	Response utils.ToolsResponse
	Usecase  usecases.UsecaseUploads
}

func New() *HttHandlersUpload {

	databases.Initpostgre(context.Background())
	bulkConnection := databases.InitBulkConnectionPkg()

	repoUpload := repositories.NewUploadBulkConnectionRepo()
	usecaseUpload := usecases.NewUsecaseUploads(repoUpload, *bulkConnection)

	return &HttHandlersUpload{
		Logger:   utils.NewToolsLogger(),
		Response: utils.NewToolsReponse(),
		Usecase:  usecaseUpload,
	}

}

func (h *HttHandlersUpload) Mount(echoGroup *echo.Group) {
	echoGroup.POST("/", h.Upload)
}

func (h *HttHandlersUpload) Upload(c echo.Context) error {
	defer h.Logger.LoggerError(c)

	fileHeader, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}

	file, err := fileHeader.Open()
	if err != nil {
		panic(err)
	}

	payload := models.UploadRequest{
		ID:        c.FormValue("id"),
		TableName: c.FormValue("table_name"),
		File:      file,
	}

	res := h.Usecase.UploadFile(c.Request().Context(), &payload)

	if res.Error != nil {
		panic(res.Error)
	}

	return h.Response.SenderResponseJSON(c, 200, "Berhasil Upload Excel", res.Data)
}
