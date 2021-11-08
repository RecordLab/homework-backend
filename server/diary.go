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

func (s *Server) GetAllDiaries(c echo.Context) error {
	diaries, err := s.ds.DiariesByUserID(c.Request().Context(), s.GetUserID(c))
	if err != nil {
		return err
	}
	type Diary struct {
		Content  string    `json:"content"`
		Image    string    `json:"image"`
		Date     time.Time `json:"date"`
		Emotions []string  `json:"emotions"`
		Theme    string    `json:"theme"`
	}
	resp := struct {
		Diaries []Diary `json:"diaries"`
	}{
		Diaries: []Diary{},
	}
	for _, diary := range diaries {
		resp.Diaries = append(resp.Diaries, Diary{
			Content:  diary.Content,
			Image:    diary.Image,
			Date:     diary.Date,
			Emotions: diary.Emotions,
			Theme:    diary.Theme,
		})
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) GetCalendar(c echo.Context) error {
	var req struct {
		Date string `query:"date"`
		Type string `query:"type"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	if req.Date == "" || req.Type == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing parameter")
	}
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return err
	}
	var diaries []model.Diary
	if req.Type == "monthly" || req.Type == "weekly" {
		diaries, err = s.ds.Calendar(c.Request().Context(), s.GetUserID(c), req.Type, date)
		if err != nil {
			return err
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid type")
	}
	type Diary struct {
		Content  string    `json:"content"`
		Image    string    `json:"image"`
		Date     time.Time `json:"date"`
		Emotions []string  `json:"emotions"`
		Theme    string    `json:"theme"`
	}
	resp := struct {
		Diaries []Diary `json:"diaries"`
	}{
		Diaries: []Diary{},
	}
	for _, diary := range diaries {
		resp.Diaries = append(resp.Diaries, Diary{
			Content:  diary.Content,
			Image:    diary.Image,
			Date:     diary.Date,
			Emotions: diary.Emotions,
			Theme:    diary.Theme,
		})
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) GetDiary(c echo.Context) error {
	dateString := c.Param("date")
	userID := s.GetUserID(c)
	date, err := time.Parse("2006-01-02", dateString)
	if dateString == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing date parameter")
	}
	if err != nil {
		return err
	}
	diary, err := s.ds.DiaryByUserIDAndDate(c.Request().Context(), userID, date)
	type Diary struct {
		Content  string    `json:"content"`
		Image    string    `json:"image"`
		Date     time.Time `json:"date"`
		Emotions []string  `json:"emotions"`
		Theme    string    `json:"theme"`
	}
	resp := Diary{
		Content:  diary.Content,
		Image:    diary.Image,
		Date:     diary.Date,
		Emotions: diary.Emotions,
		Theme:    diary.Theme,
	}
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "diary not found for given date")
		}
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) CreateDiary(c echo.Context) error {
	var req struct {
		Content  string
		Image    string
		Emotions []string
		Theme    string
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	theme := req.Theme
	isExists, err := s.ds.ThemeExists(c.Request().Context(), theme)
	if err != nil {
		return err
	}
	if !isExists {
		return echo.NewHTTPError(http.StatusBadRequest, "theme does not exist")
	}
	diary := model.Diary{
		Content:  req.Content,
		Image:    req.Image,
		Emotions: req.Emotions,
		UserID:   s.GetUserID(c),
		Date:     time.Now(),
		Theme:    theme,
	}
	if err := s.ds.WriteDiary(c.Request().Context(), diary); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (s *Server) DeleteDiary(c echo.Context) error {
	dateString := c.Param("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return err
	}
	if err := s.ds.DeleteDiary(c.Request().Context(), s.GetUserID(c), date); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
