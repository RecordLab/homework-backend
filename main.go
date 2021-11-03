package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"dailyscoop-backend/config"
	"dailyscoop-backend/server"
	"dailyscoop-backend/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	mc, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.Mongo.URL))
	if err != nil {
		panic(err)
	}
	defer mc.Disconnect(context.Background())

	us := service.NewUserService(cfg.Mongo, mc)
	ds := service.NewDiaryService(cfg.Mongo, mc)
	s := server.NewServer(cfg, us, ds)

	s.RegisterRoutes()

	s.Logger.Fatal(s.Start(cfg.Server.BindAddr))
}
