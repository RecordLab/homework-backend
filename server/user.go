package server

import (
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/labstack/echo/v4"
)

func (s *Server) Login(c echo.Context) error {
	var req struct {
		Username string
		Password string
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	user, err := s.us.UserByUsername(c.Request().Context(), req.Username)
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
