package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"dailyscoop-backend/config"
	"dailyscoop-backend/model"
)

type UserService struct {
	cfg config.MongoConfig
	mc  *mongo.Client
}

func NewUserService(cfg config.MongoConfig, mc *mongo.Client) *UserService {
	return &UserService{
		cfg: cfg,
		mc:  mc,
	}
}

func (us *UserService) UserByID(ctx context.Context, id string) (model.User, error) {
	coll := us.mc.Database(us.cfg.Database).Collection("users")
	var user model.User
	if err := coll.FindOne(ctx, bson.M{model.UserIDKey: id}).Decode(&user); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (us *UserService) UserByNickname(ctx context.Context, nickname string) (model.User, error) {
	coll := us.mc.Database(us.cfg.Database).Collection("users")
	var user model.User
	if err := coll.FindOne(ctx, bson.M{model.UserNicknameKey: nickname}).Decode(&user); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (us *UserService) RegisterUser(ctx context.Context, user model.User) error {
	coll := us.mc.Database(us.cfg.Database).Collection("users")
	h, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(h)
	if _, err := coll.InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}
