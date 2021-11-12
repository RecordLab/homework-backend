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
	ds  *service.DiaryService
	as  *service.AWSService
}

func NewServer(cfg config.Config, us *service.UserService, ds *service.DiaryService, as *service.AWSService) *Server {
	s := &Server{
		Echo: echo.New(),
		cfg:  cfg,
		us:   us,
		ds:   ds,
		as:   as,
	}
	s.Use(middleware.Logger())
	s.Use(middleware.Recover())
	return s
}

func (s *Server) RegisterRoutes() {
	api := s.Group("/api")

	api.POST("/login", s.Login)
	api.POST("/signup", s.SignUp)
	api.POST("/image", s.ImageUpload)

	diaries := api.Group("/diaries")
	diaries.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte(s.cfg.Server.Secret),
	}))
	diaries.GET("", s.GetAllDiaries)
	diaries.GET("/calendar", s.GetCalendar)
	diaries.POST("", s.CreateDiary)
	diaries.GET("/:date", s.GetDiary)
	diaries.DELETE("/:date", s.DeleteDiary)

}
