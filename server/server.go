package server

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
}

func NewServer() *Server {
	return &Server{
		Echo: echo.New(),
	}
}
