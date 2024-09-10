package main

import (
	"backend/controller"
	"backend/middlewares"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	router := mux.NewRouter()

	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/uploadImage", controller.UploadImagesHandler).Methods("POST")
	router.HandleFunc("/sendConfirmationEmail", controller.EmailConfirmationWebhook).Methods("POST")
	router.HandleFunc("/verifyEmail", controller.ConfirmEmail).Methods("GET")
	router.HandleFunc("/ticketPayment", controller.InitializePayment).Methods("POST")

	handler := loggingMiddleware(corsMiddleware(router))

	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}

	fmt.Printf("Server is running on port %s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
