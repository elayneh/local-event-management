package middlewares

import (
    "log"
    "net/http"
)

// Logger logs each request
func Logger(r *http.Request) {
    log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.RemoteAddr)
}
