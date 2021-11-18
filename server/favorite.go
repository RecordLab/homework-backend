package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) AddFavorite(c echo.Context) error {
	var req struct {
		Quote string
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	if req.Quote == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "즐겨찾기에 추가할 명언을 선택해주세요.")
	}
	userID := s.GetUserID(c)
	isExists, err := s.fs.IsFavoriteExists(c.Request().Context(), userID, req.Quote)
	if err != nil {
		return err
	}
	if isExists {
		return echo.NewHTTPError(http.StatusBadRequest, "이미 즐겨찾기에 추가되어있는 명언입니다.")
	}
	if err := s.fs.AddFavorite(c.Request().Context(), userID, req.Quote); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "즐겨찾기에 명언을 추가하였습니다.",
	})
}

func (s *Server) GetFavorites(c echo.Context) error {
	userID := s.GetUserID(c)
	favorites, err := s.fs.FavoritesByUserID(c.Request().Context(), userID)
	if err != nil {
		return err
	}
	type Favorite struct {
		Quote string `json:"quote"`
	}
	resp := struct {
		Favorite []Favorite `json:"favorite"`
	}{
		Favorite: []Favorite{},
	}
	for _, favorite := range favorites {
		resp.Favorite = append(resp.Favorite, Favorite{
			Quote: favorite.Quote,
		})
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *Server) DeleteFavorite(c echo.Context) error {
	var req struct {
		Quote string
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	userID := s.GetUserID(c)
	isExist, err := s.fs.IsFavoriteExists(c.Request().Context(), userID, req.Quote)
	if err != nil {
		return nil
	}
	if req.Quote == "" || !isExist {
		return echo.NewHTTPError(http.StatusBadRequest, "즐겨찾기를 해제할 명언이 없습니다.")
	}
	if err := s.fs.DeleteFavorite(c.Request().Context(), userID, req.Quote); err != nil {
		return nil
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "즐겨찾기를 해제했습니다.",
	})
}
