package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"

	"github.com/labstack/echo/v4"

	"dailyscoop-backend/model"
)

type jwtCustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func (s *Server) GetUserID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	id := claims.ID
	return id
}

func (s *Server) Login(c echo.Context) error {
	typ := c.QueryParam("type")
	var user model.User
	if typ == "kakao" {
		var err error
		user, err = s.KakaoLogin(c)
		if err != nil {
			return err
		}
	} else if typ == "google" {
		var err error
		user, err = s.GoogleLogin(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	} else {
		var err error
		var req struct {
			ID       string
			Password string
		}
		if err := c.Bind(&req); err != nil {
			return err
		}
		user, err = s.us.UserByID(c.Request().Context(), req.ID)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return echo.NewHTTPError(http.StatusUnauthorized, "아이디나 비밀번호를 확인해주세요.")
			}
			return err
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "아이디나 비밀번호를 확인해주세요.")
		}
	}

	claims := &jwtCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.cfg.Server.Secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token":    t,
		"nickname": user.Nickname,
	})
}

func (s *Server) KakaoLogin(c echo.Context) (model.User, error) {
	token := "Bearer " + c.Request().Header.Get("Authorization")
	req, err := http.NewRequest("GET", "https://kapi.kakao.com/v2/user/me", nil)
	if err != nil {
		return model.User{}, err
	}
	req.Header.Add("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.User{}, err
	}
	defer resp.Body.Close()
	type Response struct {
		ID          int
		ConnectedAt string `json:"connected_at"`
		Properties  struct {
			Nickname string
		}
		Message string `json:"msg"`
		Code    int
	}
	var result Response
	bytes, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(bytes, &result); err != nil {
		return model.User{}, err
	}
	if resp.StatusCode != 200 {
		return model.User{}, echo.NewHTTPError(resp.StatusCode, result.Message)
	}
	user, err := s.us.UserByID(c.Request().Context(), strconv.Itoa(result.ID))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			user = model.User{
				ID:       strconv.Itoa(result.ID),
				Nickname: result.Properties.Nickname,
			}
			if err := s.us.RegisterUser(c.Request().Context(), user); err != nil {
				return model.User{}, err
			}
		} else {
			return model.User{}, err
		}
	}
	return user, nil
}

func (s *Server) GoogleLogin(c echo.Context) (model.User, error) {
	type TokenInfo struct {
		Email   string
		Name    string
		Picture string
	}
	idToken := c.QueryParam("id_token")
	v, err := idtoken.Validate(c.Request().Context(), idToken, os.Getenv("GOOGLE_KEY"))
	if err != nil {
		return model.User{}, err
	}
	var tokenInfo TokenInfo
	if err := mapstructure.Decode(v.Claims, &tokenInfo); err != nil {
		return model.User{}, err
	}
	user, err := s.us.UserByID(c.Request().Context(), tokenInfo.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			user = model.User{
				ID:           tokenInfo.Email,
				Nickname:     tokenInfo.Name,
				ProfileImage: tokenInfo.Picture,
			}
			if err := s.us.RegisterUser(c.Request().Context(), user); err != nil {
				return model.User{}, err
			}
		} else {
			return model.User{}, err
		}
	}
	return user, nil
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
	if req.ID == "" || req.Password == "" || req.Nickname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "파라미터가 올바르지 않습니다.")
	}
	ctx := c.Request().Context()
	_, err := s.us.UserByID(ctx, req.ID)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	} else if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "이미 존재하는 아이디입니다.")
	}

	_, err = s.us.UserByNickname(ctx, req.Nickname)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	} else if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "이미 존재하는 닉네임입니다.")
	}

	if err := s.us.RegisterUser(ctx, model.User{
		ID:       req.ID,
		Password: req.Password,
		Nickname: req.Nickname,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "회원가입이 완료되었습니다.",
	})
}

func (s *Server) DeleteUser(c echo.Context) error {
	if err := s.us.DeleteUser(c.Request().Context(), s.GetUserID(c)); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusBadRequest, "존재하지 않는 유저입니다.")
		}
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "회원탈퇴가 완료되었습니다.",
	})
}

func (s *Server) ChangeNickname(c echo.Context) error {
	var req struct {
		NewNickname string
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	if req.NewNickname == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "변경할 닉네임을 입력해주세요.")
	}
	if err := s.us.UpdateNickname(c.Request().Context(), s.GetUserID(c), req.NewNickname); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "닉네임을 변경했습니다.",
	})
}

func (s *Server) ChangePassword(c echo.Context) error {
	var req struct {
		Password    string
		NewPassword string
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	if req.Password == "" || req.NewPassword == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "비밀번호와 새 비밀번호를 모두 입력해주세요.")
	}
	userID := s.GetUserID(c)
	user, err := s.us.UserByID(c.Request().Context(), userID)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "비밀번호가 일치하지 않습니다.")
	}
	if err := s.us.UpdatePassword(c.Request().Context(), userID, req.NewPassword); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "비밀번호가 변경되었습니다.",
	})
}

func (s *Server) GetUserInfo(c echo.Context) error {
	user, err := s.us.UserByID(c.Request().Context(), s.GetUserID(c))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusBadRequest, "존재하지 않는 유저입니다.")
		}
		return err
	}
	type resp struct {
		ID           string `json:"id"`
		Nickname     string `json:"nickname"`
		ProfileImage string `json:"profile_image"`
	}
	return c.JSON(http.StatusOK, resp{
		ID:           user.ID,
		Nickname:     user.Nickname,
		ProfileImage: user.ProfileImage,
	})
}

func (s *Server) SetProfileImage(c echo.Context) error {
	var req struct {
		Image string
	}
	if err := c.Bind(&req); err != nil {
		return err
	}
	if req.Image == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "사진을 선택해주세요.")
	}
	if err := s.us.UpdateProfileImage(c.Request().Context(), s.GetUserID(c), req.Image); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "프로필 사진을 변경하였습니다.",
	})
}
