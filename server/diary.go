package server

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"dailyscoop-backend/model"
)

func (s *Server) GetDiaries(c echo.Context) error {
	var req struct {
		UserID string
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	diaries, err := s.ds.DiariesByUserID(c.Request().Context(), req.UserID)
	if err != nil {
		return err
	}
	if diaries == nil {
		diaries = []model.Diary{}
	}
	return c.JSON(http.StatusOK, struct {
		Data []model.Diary `json:"data"`
	}{
		Data: diaries,
	})
}