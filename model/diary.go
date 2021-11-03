package model

import (
	"time"
)

const (
	DiaryContentKey  = "content"
	DiaryImageKey    = "image"
	DiaryUserIDKey   = "user_id"
	DiaryDateKey     = "date"
	DiaryEmotionsKey = "emotions"
)

type Diary struct {
	Content  string
	Image    string
	UserID   string `bson:"user_id"`
	Date     time.Time
	Emotions []string
}
