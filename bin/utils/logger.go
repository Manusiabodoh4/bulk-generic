package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ToolsLogger interface {
	LoggerError(c echo.Context)
}

type ToolsLoggerImpl struct {
	Response ToolsResponse
}

func NewToolsLogger() ToolsLogger {
	return &ToolsLoggerImpl{Response: NewToolsReponse()}
}

func (st *ToolsLoggerImpl) LoggerError(c echo.Context) {
	message := recover()
	if message != nil {
		st.Response.SenderResponseJSON(c, http.StatusInternalServerError, "Server Error", nil)
		c.Logger().Error(message)
	}
}
