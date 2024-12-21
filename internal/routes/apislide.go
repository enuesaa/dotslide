package routes

import (
	"os"

	"github.com/enuesaa/dotslide/internal/slides"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
)

func HandleApiSlide(c echo.Context) error {
	data, err := os.ReadFile("testdata/.slide.yml")
	if err != nil {
		return err
	}
	
	var dotslide slides.DotSlide
	if err := yaml.Unmarshal(data, &dotslide); err != nil {
		return err
	}

	return c.JSON(200, dotslide)
}
