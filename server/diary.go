package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"

	"dailyscoop-backend/model"
)

func (s *Server) GetUserID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	id := claims.ID
	return id
}

func (s *Server) GetDiaries(c echo.Context) error {
	diaries, err := s.ds.DiariesByUserID(c.Request().Context(), s.GetUserID(c))
	if err != nil {
		return err
	}
	type Diary struct {
		Content  string    `json:"content"`
		Image    string    `json:"image"`
		Date     time.Time `json:"date"`
		Emotions []string  `json:"emotions"`
	}
	resp := []Diary{}
	for _, diary := range diaries {
		resp = append(resp, Diary{
			Content:  diary.Content,
			Image:    diary.Image,
			Date:     diary.Date,
			Emotions: diary.Emotions,
		})
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) GetDiary(c echo.Context) error {
	dateString := c.Param("date")
	userID := s.GetUserID(c)
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return err
	}
	diary, err := s.ds.DiaryByUserIDAndDate(c.Request().Context(), userID, date)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "diary not found for given date")
		}
		return err
	}
	return c.JSON(http.StatusOK, diary)
}

func (s *Server) CreateDiary(c echo.Context) error {
	var req struct {
		Content  string
		Image    string
		Emotions []string
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	diary := model.Diary{
		Content:  req.Content,
		Image:    req.Image,
		Emotions: req.Emotions,
		UserID:   s.GetUserID(c),
		Date:     time.Now(),
	}
	if err := s.ds.WriteDiary(c.Request().Context(), diary); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
