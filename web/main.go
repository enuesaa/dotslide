package web

import (
	"github.com/enuesaa/dotslide/web/routes"
	"github.com/labstack/echo/v4"
)

func New() {
	app := echo.New()

	app.GET("/api/slide", routes.HandleApiSlide)
	app.GET("/storage/*", routes.HandleStorage)
	app.GET("/*", routes.HandleUi)

	app.Logger.Fatal(app.Start(":3000"))
}
