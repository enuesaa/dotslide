package web

import (
	"github.com/enuesaa/dotslide/web/routes"
	"github.com/labstack/echo/v4"
)

func NewReouter() *echo.Echo {
	app := echo.New()

	app.GET("/api/slide", routes.HandleApiSlide)
	app.GET("/storage/*", routes.HandleStorage)
	app.GET("/*", routes.HandleUi)

	return app
}
