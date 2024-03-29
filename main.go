package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/urizennnn/GO-PROJECTS/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	router := chi.NewRouter()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dbURL := os.Getenv("DBURL")
	if dbURL == "" {
		log.Println("URL not found")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in .env")
	}

	connect, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	queries := database.New(connect)

	apiCfg := apiConfig{
		DB: queries,
	}

	log.Printf("Server starting on port %v", portString)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	router.Mount("/v1", v1router)

	v1router.Post("/users", apiCfg.handleUser)
	v1router.Get("/ready", handlerReadines)
	v1router.Get("/err", handlererr)
	v1router.Get("/users", apiCfg.middlwareAuth(apiCfg.getUserbyApi))
	v1router.Post("/feed", apiCfg.middlwareAuth(apiCfg.handlerCreateFeed))
	v1router.Get("/feed", apiCfg.getFeed)
	v1router.Post("/follows_feed",apiCfg.middlwareAuth(apiCfg.handlerCreateFeedfollows))
	v1router.Post("/follows_feed",apiCfg.middlwareAuth(apiCfg.handlerCreateFeedfollows))
	v1router.Get("/follows_feed",apiCfg.middlwareAuth(apiCfg.GetUserFeed))
	v1router.Delete("/follows_feed/{feedID}",apiCfg.middlwareAuth(apiCfg.DeleteFeed))

	server.ListenAndServe()
}
