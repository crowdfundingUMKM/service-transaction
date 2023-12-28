package main

import (
	"fmt"
	"log"
	"os"
	"service-transaction/auth"
	"service-transaction/config"
	"service-transaction/core"
	"service-transaction/database"
	"service-transaction/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup log if Production Mode
	// config.InitLog()
	// SETUP REPO
	db := database.NewConnectionDB()
	transactionRepository := core.NewRepository(db)

	// SETUP SERVICE {use auth api investor}
	transactionService := core.NewService(transactionRepository)
	authService := auth.NewService()

	// setup handler
	transactionHandler := handler.NewTransactionHandler(transactionService, authService)
	// END SETUP

	// RUN SERVICE
	router := gin.Default()
	// setup cors
	corsConfig := config.InitCors()
	router.Use(cors.New(corsConfig))

	// group api
	api := router.Group("api/v1")

	// uri for search campaign id and auth bearer token investor
	api.POST("/create-transaction/:campaign_id", transactionHandler.CreateTransaction)

	// api.POST("/transactions/notification", transactionHandler.GetNotification)

	// setup route

	// end Rounting
	// make env SERVICE_HOST and SERVICE_PORT
	url := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	router.Run(url)
}
