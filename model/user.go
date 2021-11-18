package model

const (
	UserIDKey           = "id"
	UserNicknameKey     = "nickname"
	UserPasswordKey     = "password"
	UserProfileImageKey = "profile_image"
)

type User struct {
	ID           string
	Password     string
	Nickname     string
	ProfileImage string `bson:"profile_image"`
}
