package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"dailyscoop-backend/server"
	"dailyscoop-backend/service"
)

func main() {
	mc, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		panic(err)
	}
	defer mc.Disconnect(context.Background())

	us := service.NewUserService(mc)
	s := server.NewServer(us)

	s.RegisterRoutes()

	s.Logger.Fatal(s.Start(":8080"))
}