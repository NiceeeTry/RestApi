package main

import (
	"log"
	"rest_api/pkg/handler"
	"rest_api/todo"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error in running the server: %s", err)
	}
}
