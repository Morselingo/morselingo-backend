package router

import "net/http"

func AuthRouter() http.Handler {
	mux := http.NewServeMux()
	// mux.HandleFunc("/register", )
	return mux
}
