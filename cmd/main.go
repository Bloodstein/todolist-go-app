package main

// "log"
import (
	"log"
	"os"
	"strings"

	server "github.com/Bloodstein/todolist-go-app"
	"github.com/Bloodstein/todolist-go-app/app/handler"
	"github.com/Bloodstein/todolist-go-app/app/repository"
	"github.com/Bloodstein/todolist-go-app/app/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	password := os.Getenv("DB_PASSWORD")

	connString := strings.Replace(viper.GetString("connectionString"), "PASSWORD", password, 1)

	db, err := repository.NewPostgresDB(connString)

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
