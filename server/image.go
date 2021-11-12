package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) ImageUpload(c echo.Context) error {
	file, header, err := c.Request().FormFile("file")
	if file == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "파라미터가 잘못되었습니다.")
	}
	if err != nil {
		return err
	}
	fileName := header.Filename
	url, err := s.as.UploadImage(file, fileName)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"url": url,
	})
}
