package middleware

// "net/http"

// "github.com/gorilla/handlers"

// func ApplyCORS(handler http.Handler) http.Handler {
// 	return handlers.CORS(
// 		handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Adjust if necessary
// 		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
// 		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
// 	)(handler)
// }
// package middleware

// import (
// 	"backend/utils"
// 	"log"
// 	"net/http"
// )

// func ApplyCORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
// 		if r.Method == "OPTIONS" {
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

// func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		token := r.Header.Get("Authorization")
// 		if token == "" {
// 			http.Error(w, "Missing token", http.StatusUnauthorized)
// 			return
// 		}

// 		claims, err := utils.ValidateJWT(token)
// 		if err != nil {
// 			http.Error(w, "Invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		log.Printf("Authenticated user: %s", claims.Username)

// 		next(w, r)
// 	}
// }
