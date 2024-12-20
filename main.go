package main

import (
	"embed"
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

//go:embed all:dist/*
var dist embed.FS

func main() {
	app := echo.New()

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
