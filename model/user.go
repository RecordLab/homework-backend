package model

const (
	UserIDKey       = "id"
	UserPasswordKey = "password"
	UserNicknameKey = "nickname"
)

type User struct {
	ID       string
	Password string
	Nickname string
}
