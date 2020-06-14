package image

import (
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/nauhalf/mobilcantik/utils"
)

func StreamImage(c echo.Context) error {
	imageDir := utils.GetImageDir()
	imagePath := imageDir + c.Param("*")
	imageFile, err := os.Open(imagePath)
	if err != nil {
		return c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	defer imageFile.Close()

	return c.Stream(200, mime.TypeByExtension(filepath.Ext(imageFile.Name())), imageFile)
}
