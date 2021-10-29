package main

import (
	"github.com/dailyscoop/dailyscoop-backend/server"
)

func main() {
	s := server.NewServer()

	s.GET("/login", s.Login)

	s.Logger.Fatal(s.Echo.Start(":8080"))
}