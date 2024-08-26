package main

import (
	"backend/controller"
	"backend/middlewares"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middlewares.Logger(r)
		next.ServeHTTP(w, r)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middlewares.CorsMiddleware(w, r)
		next.ServeHTTP(w, r)
	})
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/register", controller.Register)
	mux.HandleFunc("/login", controller.Login)

	mux.HandleFunc("/uploadImage", controller.UploadImagesHandler)

	handler := loggingMiddleware(corsMiddleware(mux))
	port := os.Getenv("PORT")
	fmt.Println("Server is running on port 5050")

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
