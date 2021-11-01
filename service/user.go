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

func (us *UserService) UserByID(ctx context.Context, ID string) (model.User, error) {
	coll := us.mc.Database("dailyscoop").Collection("users")
	var user model.User
	if err := coll.FindOne(ctx, bson.M{model.UserIDKey: ID}).Decode(&user); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (us *UserService) UserByNickname(ctx context.Context, Nickname string) (model.User, error) {
	coll := us.mc.Database("daliyscoop").Collection("users")
	var user model.User
	if err := coll.FindOne(ctx, bson.M{model.UserNicknameKey: Nickname}).Decode(&user); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (us *UserService) RegisterUser(ctx context.Context, user model.User) error {
	coll := us.mc.Database("dailyscoop").Collection("users")
	_, err := coll.InsertOne(ctx, bson.M{
		model.UserIDKey: user.ID,
		model.UserNicknameKey: user.Nickname,
		model.UserPasswordKey: user.Password,
	})
	if err != nil {
		return err
	}
	return nil
}
