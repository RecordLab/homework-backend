package model

const (
	FavoriteUserIDKey  = "user_id"
	FavoriteContentKey = "quote"
)

type Favorite struct {
	UserID string `bson:"user_id"`
	Quote  string
}
