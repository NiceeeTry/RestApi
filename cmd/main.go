package main

import (
	"log"
	"rest_api/pkg/handler"
	"rest_api/pkg/repository"
	"rest_api/pkg/service"
	"rest_api/todo"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error in configs: %s", err.Error)
	}
	repos := repository.NewRepository()
	services := service.NewService(*repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatal("error in running the server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
