package model

const (
	UserIDKey           = "id"
	UserKakaoIDKey      = "kakao_id"
	UserNicknameKey     = "nickname"
	UserPasswordKey     = "password"
	UserProfileImageKey = "profile_image"
)

type User struct {
	ID           string
	KakaoID      int `bson:"kakao_id"`
	Password     string
	Nickname     string
	ProfileImage string `bson:"profile_image"`
}
