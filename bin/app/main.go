package main

import (
	handlersUpload "github.com/Manusiabodoh4/bulk-generic/bin/modules/upload/handlers"
	handlersUser "github.com/Manusiabodoh4/bulk-generic/bin/modules/users/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.CORS())

	usersGroup := app.Group("/v1/api/users")
	uploadGroup := app.Group("v1/api/upload")

	handlersUser.New().Mount(usersGroup)
	handlersUpload.New().Mount(uploadGroup)

	app.Logger.Fatal(app.Start(":4567"))

}
