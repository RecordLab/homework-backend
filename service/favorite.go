package service

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"dailyscoop-backend/config"
	"dailyscoop-backend/model"
)

type FavoriteService struct {
	cfg config.MongoConfig
	mc  *mongo.Client
}

func NewFavoriteService(cfg config.MongoConfig, mc *mongo.Client) *FavoriteService {
	return &FavoriteService{
		cfg: cfg,
		mc:  mc,
	}
}

func (fs *FavoriteService) AddFavorite(ctx context.Context, userID string, quote string) error {
	coll := fs.mc.Database(fs.cfg.Database).Collection("favorites")
	favorite := model.Favorite{
		UserID: userID,
		Quote:  quote,
	}
	if _, err := coll.InsertOne(ctx, favorite); err != nil {
		return err
	}
	return nil
}

func (fs *FavoriteService) FavoritesByUserID(ctx context.Context, userID string) ([]model.Favorite, error) {
	coll := fs.mc.Database(fs.cfg.Database).Collection("favorites")
	cursor, err := coll.Find(ctx, bson.M{
		model.FavoriteUserIDKey: userID,
	})
	if err != nil {
		return nil, err
	}
	var favorites []model.Favorite
	for cursor.Next(ctx) {
		var favorite model.Favorite
		if err := cursor.Decode(&favorite); err != nil {
			return nil, err
		}
		favorites = append(favorites, favorite)
	}
	return favorites, nil
}

func (fs *FavoriteService) DeleteFavorite(ctx context.Context, userID string, quote string) error {
	coll := fs.mc.Database(fs.cfg.Database).Collection("favorites")
	if _, err := coll.DeleteOne(ctx, bson.M{
		model.FavoriteUserIDKey:  userID,
		model.FavoriteContentKey: quote,
	}); err != nil {
		return err
	}
	return nil
}

func (fs *FavoriteService) IsFavoriteExists(ctx context.Context, userID string, quote string) (bool, error) {
	coll := fs.mc.Database(fs.cfg.Database).Collection("favorites")
	err := coll.FindOne(ctx, bson.M{
		model.FavoriteUserIDKey:  userID,
		model.FavoriteContentKey: quote,
	}).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, nil
	}
	return true, nil
}
