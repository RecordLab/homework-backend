package server

import (
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/labstack/echo/v4"
)

func (s *Server) Login(c echo.Context) error {
	var req struct {
		ID string
		Password string
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	user, err := s.us.UserByID(c.Request().Context(), req.ID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusUnauthorized, "wrong username or password")
		}
		return err
	}
	if user.Password != req.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "wrong username or password")
	}
	return c.JSON(http.StatusOK, struct {
		Token string `json:"token"`
	}{
		Token: "ABCDEF",
	})
}

func (s *Server) SignUp(c echo.Context) error {
	var req struct {
		ID string
		Password string
		Nickname string
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	_, err := s.us.UserByID(c.Request().Context(), req.ID)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	} else if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "이미 존재하는 ID입니다.")
	}

	_, err = s.us.UserByNickname(c.Request().Context(), req.Nickname)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	} else if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "이미 존재하는 Nickname입니다.")
	}

	if err := s.us.RegisterUser(c.Request().Context(), req); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}
