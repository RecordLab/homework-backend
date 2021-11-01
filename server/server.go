package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"dailyscoop-backend/service"
)

type Server struct {
	*echo.Echo
	us *service.UserService
}

func NewServer(us *service.UserService) *Server {
	s := &Server{
		Echo: echo.New(),
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
