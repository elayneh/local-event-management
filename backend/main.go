package main

import (
	"backend/controller"
	"backend/middlewares"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const uploadPath = "./uploads"

func init() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create the uploads directory if it doesn't exist
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err = os.Mkdir(uploadPath, 0755)
		if err != nil {
			log.Fatal("Error creating uploads directory")
		}
	}
}

func main() {
	server := gin.New()

	// Define middlewares
	server.Use(middlewares.Logger())
	server.Use(middlewares.CorsMiddleware())

	// Define routes
	server.POST("/register", controller.Register)
	server.POST("/login", controller.Login)
	server.POST("/updateUser", controller.UpdateUser)
	server.POST("/uploadFile", controller.UploadFileHandler)

	// Serve static files from the upload directory
	server.Static("/uploads", uploadPath)

	fmt.Println("Server is running on port 5050")

	// Run the server on port 5050
	server.Run(":5050")
}
