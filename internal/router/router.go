package router

import (
	"net/http"
)

func AuthRouter(registrationHandler http.HandlerFunc) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", registrationHandler)
	return mux
}

func ChatRouter() http.Handler {
	mux := http.NewServeMux()
	// mux.HandleFunc("")
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
