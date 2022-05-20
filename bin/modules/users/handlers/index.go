package handlers

import (
	"context"
	"errors"
	"fmt"

	helper "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/helper/connection"
	models "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/models/domain"
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

	fmt.Println(usecaseUsers.SetupConnection(context.Background()))

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
	defer h.Logger.LoggerError(c)

	var data models.RegisterRequest

	err := c.Bind(&data)
	if err != nil {
		panic(err)
	}

	result := h.Usecase.Register(c.Request().Context(), &data)

	if result.Error != nil {
		panic(result.Error)
	}

	return h.Reponse.SenderResponseJSON(c, 200, "Data berhasil terdaftar", result)
}

func (h *HTTPHandlers) GetDetail(c echo.Context) error {
	defer h.Logger.LoggerError(c)

	var data models.GetDetailRequest
	data.ID = c.Param("id")

	if len(data.ID) <= 0 {
		panic(errors.New("data id tidak ditemukan"))
	}

	result := h.Usecase.GetDetailWithID(c.Request().Context(), &data)

	if result.Error != nil {
		panic(result.Error)
	}

	return h.Reponse.SenderResponseJSON(c, 200, "Berhasil mendapatkan data detail", result.Data)
}

func (h *HTTPHandlers) UpdateWithID(c echo.Context) error {
	defer h.Logger.LoggerError(c)

	var data models.UpdateUsersRequest
	err := c.Bind(&data)
	if err != nil {
		panic(err)
	}

	userID := c.Param("id")
	if len(userID) <= 0 {
		panic(errors.New("data id tidak ditemukan"))
	}

	result := h.Usecase.UpdateUsers(c.Request().Context(), userID, &data)

	return h.Reponse.SenderResponseJSON(c, 200, "Berhasil melakukan update data", result.Data)
}
