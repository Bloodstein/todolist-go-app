package main

// "log"
import (
	"log"

	server "github.com/Bloodstein/todolist-go-app"
	"github.com/Bloodstein/todolist-go-app/app/handler"
	"github.com/Bloodstein/todolist-go-app/app/repository"
	"github.com/Bloodstein/todolist-go-app/app/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "kashin.db.elephantsql.com",
		Port:     "5432",
		Username: "xdbejrew",
		Password: "fYqD3Kt4xFFkD4vlASBEPd8jJGSMvmxM",
		DBName:   "xdbejrew",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatalf("fail to initializing db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
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
