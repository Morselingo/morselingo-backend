package router

import (
	"net/http"

	"github.com/Morselingo/morselingo-backend/internal/auth"
)

func AuthRouter(registrationHandler, loginHandler http.HandlerFunc) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", registrationHandler)
	mux.HandleFunc("/login", loginHandler)
	return mux
}

func ChatRouter(subscribeHandler http.HandlerFunc) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/ws", auth.JWTMiddleware(subscribeHandler))
	return mux
}

func UserRouter() http.Handler {
	mux := http.NewServeMux()
	// mux.HandleFunc("")
	return mux
}

func LeaderbordRouter() http.Handler {
	mux := http.NewServeMux()
	return mux
}
