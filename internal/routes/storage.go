package routes

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

func HandleStorage(c echo.Context) error  {
	path := c.Request().URL.Path // like `/`
	path = strings.TrimPrefix(path, "/storage/")


	fileExt := filepath.Ext(path)
	fmt.Println("testdata/" + path)
	f, err := os.ReadFile("testdata/" + path)
	if err != nil {
		return err
	}
	mimeType := mime.TypeByExtension(fileExt)

	return c.Blob(200, mimeType, f)
}
