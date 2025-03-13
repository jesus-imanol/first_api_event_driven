package main

import (
	"apiInvitation/src/match/infaestructure/configMatch"
	"apiInvitation/src/users/infraestructure/config"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Authorization"}, 
		MaxAge:           12 * time.Hour,
	}))

	config.InitUsers(r)
	configMatch.InitMatch(r)
	

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
