package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"dailyscoop-backend/model"
)

type UserService struct {
	mc *mongo.Client
}

func NewUserService(mc *mongo.Client) *UserService {
	return &UserService{
		mc: mc,
	}
}

func (us *UserService) UserByUsername(ctx context.Context, username string) (model.User, error) {
	coll := us.mc.Database("dailyscoop").Collection("users")
	var user model.User
	if err := coll.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		return model.User{}, err
	}
	return user, nil
}
