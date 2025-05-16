package router

import "net/http"

func LeaderbordRouter() http.Handler {
	mux := http.NewServeMux()
	return mux
}
