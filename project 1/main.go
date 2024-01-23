package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/urizennnn/GO-PROJECTS/internal/database"
)

func main() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1router := chi.NewRouter()
	type apiConfig struct {
		DB *database.Queries
	}

	router.Mount("/v1", v1router)

	v1router.Get("/ready", handlerReadines)
	v1router.Get("/err", handlererr)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in .env")
	}
	log.Printf("Server starting on port %v", portString)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	server.ListenAndServe()

	fmt.Println("PORT:", portString)
}
