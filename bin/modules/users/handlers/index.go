package handlers

import (
	"context"

	helper "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/helper/connection"
	"github.com/Manusiabodoh4/bulk-generic/bin/modules/users/repositories"
	"github.com/Manusiabodoh4/bulk-generic/bin/modules/users/usecases"
	"github.com/Manusiabodoh4/bulk-generic/bin/pkg/databases"
	"github.com/Manusiabodoh4/bulk-generic/bin/utils"
	"github.com/labstack/echo/v4"
)

type HTTPHandlers struct {
	Logger  utils.ToolsLogger
	Reponse utils.ToolsResponse
	Usecase usecases.UsecaseUsers
}

func New() *HTTPHandlers {

	helper.NewCreateConnectionHelper()

	postgresDB := databases.Initpostgre(context.Background())

	bulkConnection := databases.InitBulkConnectionPkg()
	bulkConnection.AddBulkConnection("main", postgresDB)

	repoUsers := repositories.NewUserPostgresImpl(postgresDB)
	usecaseUsers := usecases.NewUsecaseUsers(repoUsers, *bulkConnection)

	return &HTTPHandlers{
		Logger:  utils.NewToolsLogger(),
		Reponse: utils.NewToolsReponse(),
		Usecase: usecaseUsers,
	}
}

func (h *HTTPHandlers) Mount(echoGroup *echo.Group) {
	echoGroup.POST("/register", h.Register)
	echoGroup.GET("/:id", h.GetDetail)
	echoGroup.PUT("/:id", h.UpdateWithID)
}

func (h *HTTPHandlers) Register(c echo.Context) error {

	return nil
}

func (h *HTTPHandlers) GetDetail(c echo.Context) error {
	return nil
}

func (h *HTTPHandlers) UpdateWithID(c echo.Context) error {
	return nil
}
