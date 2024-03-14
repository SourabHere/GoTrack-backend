package main

import (
	"fmt"

	"example.com/app"
	"example.com/db"
	"example.com/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	db, err := db.NewPostgresDB()

	if err != nil {
		panic(err)
	}

	handlers := app.IntitialiseHandlers(db.DB)

	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	server.Use(cors.New(config))

	routes.RegisterRoutes(server, handlers)

	server.Run(":8080")

}
