package routes

import (
	"os"

	"github.com/enuesaa/dotslide/internal/slides"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
)

func HandleApiSlide(c echo.Context) error {
	data, err := os.ReadFile("slides.yml")
	if err != nil {
		return err
	}

	type Response struct {
		Slides []slides.Slide `yaml:"slides" json:"slides"`
	}
	
	var response Response
	err = yaml.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	return c.JSON(200, response)
}
