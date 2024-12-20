package main

import (
	"embed"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
)

//go:embed all:dist/*
var dist embed.FS

func main() {
	app := echo.New()

	app.GET("/api/slide", func(c echo.Context) error {
		data, err := os.ReadFile("slides.yml")
		if err != nil {
			return err
		}

		type Response struct {
			Slides []Slide `yaml:"slides" json:"slides"`
		}
		
		var response Response
		err = yaml.Unmarshal(data, &response)
		if err != nil {
			return err
		}

		return c.JSON(200, response)
	})

	app.GET("/storage/*", func (c echo.Context) error  {
		path := c.Request().URL.Path // like `/`
		path = strings.TrimPrefix(path, "/storage/")
	
		fileExt := filepath.Ext(path)
		f, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		mimeType := mime.TypeByExtension(fileExt)
	
		return c.Blob(200, mimeType, f)
	})

	app.GET("/*", func(c echo.Context) error {
		path := c.Request().URL.Path // like `/`
		path = fmt.Sprintf("dist%s", path)
	
		fileExt := filepath.Ext(path)
		if fileExt == "" || strings.HasSuffix(path, "/") {
			path = "dist/index.html"
		}
	
		f, err := dist.ReadFile(path)
		if err != nil {
			return err
		}
		mimeType := mime.TypeByExtension(fileExt)
	
		return c.Blob(200, mimeType, f)
	})

	app.Logger.Fatal(app.Start(":3000"))
}
