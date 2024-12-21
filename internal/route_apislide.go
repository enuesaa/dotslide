package internal

import (
	"os"
	"path/filepath"

	"github.com/enuesaa/dotslide/internal/slides"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
)

func (r *Router) handleApiSlide(c echo.Context) error {
	path := filepath.Join(r.config.Workdir, ".slide.yml")

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	
	var dotslide slides.DotSlide
	if err := yaml.Unmarshal(data, &dotslide); err != nil {
		return err
	}

	return c.JSON(200, dotslide)
}
