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

func ChatRouter(sendMessageHandler, subscribeHandler http.HandlerFunc) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/send", auth.JWTMiddleware(sendMessageHandler))
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
