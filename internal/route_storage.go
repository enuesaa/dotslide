package internal

import (
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

func (r *Router) handleStorage(c echo.Context) error {
	path := c.Request().URL.Path // like `/`
	path = strings.TrimPrefix(path, "/storage/")
	path = filepath.Join(r.config.Workdir, path)

	fileExt := filepath.Ext(path)
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	mimeType := mime.TypeByExtension(fileExt)

	return c.Blob(200, mimeType, f)
}
