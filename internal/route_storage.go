package internal

import (
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

func (r *Router) handleStorage(c echo.Context) error {
	path := c.Request().URL.Path
	path = strings.TrimPrefix(path, "/storage/")
	path = filepath.Join(r.config.Workdir, path)

	return c.File(path)
}
