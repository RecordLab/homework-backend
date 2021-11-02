package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"dailyscoop-backend/config"
	"dailyscoop-backend/service"
)

type Server struct {
	*echo.Echo
	cfg config.Config
	us  *service.UserService
}

func NewServer(cfg config.Config, us *service.UserService) *Server {
	s := &Server{
		Echo: echo.New(),
		cfg:  cfg,
		us:   us,
	}
	s.Use(middleware.Logger())
	s.Use(middleware.Recover())
	return s
}

func (s *Server) RegisterRoutes() {
	s.POST("/login", s.Login)
	s.POST("/signup", s.SignUp)
}
