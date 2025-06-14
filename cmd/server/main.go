package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Morselingo/morselingo-backend/internal/auth"
	"github.com/Morselingo/morselingo-backend/internal/handler"
	"github.com/Morselingo/morselingo-backend/internal/repository"
	"github.com/Morselingo/morselingo-backend/internal/router"
	"github.com/Morselingo/morselingo-backend/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func main() {
	if err := auth.InitializeAuthentication(os.Getenv("JWT_SECRET_KEY")); err != nil {
		log.Fatal("Failed to initialize JWT Secret Key")
	}

	postgres := loadPostgres()
	// valkey := loadValkey()

	// repository
	userRepository := repository.NewUserRepository(postgres)

	//  service
	userService := service.NewUserService(userRepository)
	chatService := service.NewChatService()

	// handler
	userHandler := handler.NewUserHandler(userService)
	chatHandler := handler.NewChatHandler(chatService)

	// router
	authRouter := router.AuthRouter(userHandler.RegisterUser, userHandler.LoginUser)
	chatRouter := router.ChatRouter(chatHandler.Subscribe)
	leaderbordRouter := router.LeaderbordRouter()

	rootMux := http.NewServeMux()
	rootMux.Handle("/auth/", http.StripPrefix("/auth", authRouter))
	rootMux.Handle("/chat/", http.StripPrefix("/chat", chatRouter))
	rootMux.Handle("/leaderbord/", http.StripPrefix("/leaderbord", leaderbordRouter))

	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", rootMux)
	if err != nil {
		log.Fatal("Failed to create server: ", err)
	}
}

func loadPostgres() *pgxpool.Pool {
	dbURL := os.Getenv("DB_URL")
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Failed to load postgres: ", err)
	}
	return pool
}

func loadValkey() *redis.Client {
	valkeyAddr := os.Getenv("REDIS_URL")

	opts, err := redis.ParseURL(valkeyAddr)
	if err != nil {
		log.Fatal("Failed to parse REDIS_URL: ", err)
	}
	valkey := redis.NewClient(opts)

	if _, err := valkey.Ping(context.Background()).Result(); err != nil {
		log.Fatal("Failed to connect to Valkey: ", err)
	}

	return valkey
}
