package routes

import (
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/enuesaa/dotslide/internal/ui"
	"github.com/labstack/echo/v4"
)

func HandleUi(c echo.Context) error {
	path := c.Request().URL.Path // like `/`
	path = fmt.Sprintf("dist%s", path)

	fileExt := filepath.Ext(path)
	if fileExt == "" || strings.HasSuffix(path, "/") {
		path = "dist/index.html"
	}

	f, err := ui.Dist.ReadFile(path)
	if err != nil {
		return err
	}
	mimeType := mime.TypeByExtension(fileExt)

	return c.Blob(200, mimeType, f)
}
