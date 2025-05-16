package main

import (
	"log"
	"net/http"

	"github.com/Morselingo/morselingo-backend/internal/router"
)

func main() {
	rootMux := http.NewServeMux()
	rootMux.Handle("/auth/", http.StripPrefix("/auth", router.AuthRouter()))
	rootMux.Handle("/chat/", http.StripPrefix("/chat", router.ChatRouter()))
	rootMux.Handle("/leaderbord/", http.StripPrefix("/leaderbord", router.LeaderbordRouter()))

	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", rootMux)
	if err != nil {
		log.Fatal("Failed to create server: ", err)
	}
}
