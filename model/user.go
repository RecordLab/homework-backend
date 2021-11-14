package model

const (
	UserIDKey       = "id"
	UserNicknameKey = "nickname"
	UserPasswordKey = "password"
)

type User struct {
	ID           string
	Password     string
	Nickname     string
	ProfileImage string
}
