package service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"dailyscoop-backend/config"
	"dailyscoop-backend/model"
)

type DiaryService struct {
	cfg config.MongoConfig
	mc  *mongo.Client
}

func NewDiaryService(cfg config.MongoConfig, mc *mongo.Client) *DiaryService {
	return &DiaryService{
		cfg: cfg,
		mc:  mc,
	}
}

func (ds *DiaryService) DiariesByUserID(ctx context.Context, userID string) ([]model.Diary, error) {
	coll := ds.mc.Database(ds.cfg.Database).Collection("diaries")
	cursor, err := coll.Find(ctx, bson.M{model.DiaryUserIDKey: userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var diaries []model.Diary
	for cursor.Next(ctx) {
		var diary model.Diary
		if err := cursor.Decode(&diary); err != nil {
			return nil, err
		}
		diaries = append(diaries, diary)
	}
	return diaries, nil
}

func (ds *DiaryService) DiaryByUserIDAndDate(ctx context.Context, userID string, date time.Time) (model.Diary, error) {
	coll := ds.mc.Database(ds.cfg.Database).Collection("diaries")
	newDate := time.Date(
		date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	var diary model.Diary
	if err := coll.FindOne(ctx, bson.M{
		model.DiaryUserIDKey: userID,
		model.DiaryDateKey: bson.M{
			"$gte": newDate,
			"$lt":  newDate.AddDate(0, 0, 1),
		},
	}).Decode(&diary); err != nil {
		return model.Diary{}, err
	}
	return diary, nil
}

func (ds *DiaryService) WriteDiary(ctx context.Context, diary model.Diary) error {
	coll := ds.mc.Database(ds.cfg.Database).Collection("diaries")
	date := time.Date(diary.Date.Year(), diary.Date.Month(), diary.Date.Day(), 0, 0, 0, 0, diary.Date.Location())
	if _, err := coll.UpdateOne(ctx, bson.M{
		model.DiaryDateKey: bson.M{
			"$gte": date,
			"$lt":  date.AddDate(0, 0, 1),
		},
		model.DiaryUserIDKey: diary.UserID,
	}, bson.M{
		"$set": bson.M{
			model.DiaryContentKey:  diary.Content,
			model.DiaryImageKey:    diary.Image,
			model.DiaryEmotionsKey: diary.Emotions,
		},
		"$setOnInsert": bson.M{
			model.DiaryDateKey: diary.Date,
		},
	}, options.Update().SetUpsert(true)); err != nil {
		return err
	}
	return nil
}

func (ds *DiaryService) DeleteDiary(ctx context.Context, userID string, date time.Time) error {
	coll := ds.mc.Database(ds.cfg.Database).Collection("diaries")
	newDate := time.Date(
		date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	_, err := coll.DeleteOne(ctx, bson.M{
		model.DiaryUserIDKey: userID,
		model.DiaryDateKey: bson.M{
			"$gte": newDate,
			"$lt":  newDate.AddDate(0, 0, 1),
		},
	})
	if err != nil {
		return err
	}
	return nil
}
