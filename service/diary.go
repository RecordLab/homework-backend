package service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

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
	var diary model.Diary
	if err := coll.FindOne(ctx, bson.M{
		model.DiaryDateKey: date,
		model.DiaryUserIDKey: userID,
	}).Decode(&diary); err != nil {
		return model.Diary{}, err
	}
	return diary, nil
}

func (ds *DiaryService) CreateDiary(ctx context.Context, diary model.Diary) error {
	coll := ds.mc.Database(ds.cfg.Database).Collection("diaries")
	if _, err := coll.InsertOne(ctx, diary); err != nil {
		return err
	}
	return nil
}
