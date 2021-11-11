package model

const (
	UserIDKey       = "id"
	UserNicknameKey = "nickname"
)

type User struct {
	ID           string
	Password     string
	Nickname     string
	ProfileImage string
}
