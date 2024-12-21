package internal

import (
	"net/http"

	"github.com/enuesaa/dotslide/internal/middleware"
	"github.com/labstack/echo/v4"
)

func NewRouter(config Config) Router {
	return Router{
		config: config,
	}
}

type Router struct {
	config Config
}

func (r *Router) Handle() http.Handler {
	app := echo.New()

	app.Use(middleware.Logger)
	app.Use(middleware.Cors)

	app.GET("/api/dotslide", r.handleApiSlide)
	app.GET("/storage/*", r.handleStorage)
	app.GET("/*", r.handleUi)

	return app.Server.Handler
}
