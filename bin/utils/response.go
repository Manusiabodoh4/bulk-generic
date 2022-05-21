package utils

import (
	"github.com/Manusiabodoh4/bulk-generic/bin/entity"
	"github.com/labstack/echo/v4"
)

type ToolsResponse interface {
	SenderResponseJSON(c echo.Context, code int, msg interface{}, data interface{}) error
}

type ToolsResponseImpl struct{}

func NewToolsReponse() ToolsResponse {
	return &ToolsResponseImpl{}
}

func (st *ToolsResponseImpl) SenderResponseJSON(c echo.Context, code int, msg interface{}, data interface{}) error {
	res := entity.TemplateResponse{}
	res.Status = code
	res.Message = msg
	res.Data = data
	return c.JSON(code, res)
}
