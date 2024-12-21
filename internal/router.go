package internal

import (
	"net/http"

	"github.com/enuesaa/dotslide/internal/routes"
	"github.com/labstack/echo/v4"
)

func NewReouter() http.ServeMux {
	app := echo.New()

	app.GET("/api/slide", routes.HandleApiSlide)
	app.GET("/storage/*", routes.HandleStorage)
	app.GET("/*", routes.HandleUi)

	return &app.
}
