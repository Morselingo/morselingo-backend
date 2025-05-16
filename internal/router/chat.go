package router

import "net/http"

func ChatRouter() http.Handler {
	mux := http.NewServeMux()
	// mux.HandleFunc("/send", )
	return mux
}
