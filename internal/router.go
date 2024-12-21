package internal

import (
	"net/http"

	"github.com/enuesaa/dotslide/internal/middleware"
	"github.com/enuesaa/dotslide/internal/routes"
	"github.com/labstack/echo/v4"
)

func NewRouter() http.Handler {
	app := echo.New()

	app.Use(middleware.Logger)
	app.Use(middleware.Cors)

	app.GET("/api/dotslide", routes.HandleApiSlide)
	app.GET("/storage/*", routes.HandleStorage)
	app.GET("/*", routes.HandleUi)

	return app.Server.Handler
}
