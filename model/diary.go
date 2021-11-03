package model

import (
	"time"
)

const (
	DiaryContentKey  = "content"
	DiaryImageKey    = "image"
	DiaryUserIDKey   = "userID"
	DiaryDateKey     = "date"
	DiaryEmotionsKey = "emotions"
)

type Diary struct {
	Content  string
	Image    string
	UserID   string
	Date     time.Time
	Emotions []string
}
