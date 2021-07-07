package main

// "log"
import (
	"log"

	server "github.com/Bloodstein/todolist-go-app"
	"github.com/Bloodstein/todolist-go-app/app/handlers"
)

func main() {
	handlers := new(handlers.Handler)
	server := new(server.Server)
	if error := server.Run("8080", handlers.InitRoutes()); error != nil {
		log.Fatalf("Error occured while running http server: %s", error.Error())
	}
}
