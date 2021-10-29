package server

import (
	"github.com/labstack/echo/v4"

	"dailyscoop-backend/service"
)

type Server struct {
	*echo.Echo
	us *service.UserService
}

func NewServer(us *service.UserService) *Server {
	return &Server{
		Echo: echo.New(),
		us:   us,
	}
}
