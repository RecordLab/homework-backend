package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func (s *Server) Login(c echo.Context) error {
	var req struct {
		ID       string
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
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "wrong username or password")
	}

	claims := &jwtCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (s *Server) SignUp(c echo.Context) error {
	var req struct {
		ID       string
		Password string
		Nickname string
	}
	if err := c.Bind(&req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	_, err := s.us.UserByID(ctx, req.ID)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	} else if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "이미 존재하는 ID입니다.")
	}

	_, err = s.us.UserByNickname(ctx, req.Nickname)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	} else if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "이미 존재하는 Nickname입니다.")
	}

	if err := s.us.RegisterUser(ctx, req); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}
