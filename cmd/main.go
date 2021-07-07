package main

// "log"
import (
	"log"

	server "github.com/Bloodstein/todolist-go-app"
	"github.com/Bloodstein/todolist-go-app/app/handler"
	"github.com/Bloodstein/todolist-go-app/app/repository"
	"github.com/Bloodstein/todolist-go-app/app/service"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(server.Server)
	if error := server.Run(viper.GetString("port"), handlers.InitRoutes()); error != nil {
		log.Fatalf("Error occured while running http server: %s", error.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
