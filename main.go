package main

import (
	"log"
	"net/http"
	"os"

	config "github.com/Triyaambak/RSS-Aggregator/config"
	handler "github.com/Triyaambak/RSS-Aggregator/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")
	serverPort := os.Getenv("API_PORT")
	dbUrl := os.Getenv("DB_URL")

	apiCfg := config.ConnectDB(dbUrl)
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handler.HandlerReadiness)
	v1Router.Get("/err", handler.HandlerErr)
	v1Router.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		handler.HandlerCreateUser(apiCfg, w, r)
	})
	v1Router.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		handler.HandlerGetUser(apiCfg, w, r)
	})
	v1Router.Post("/feeds", func(w http.ResponseWriter, r *http.Request) {
		handler.HandlerCreateUserFeed(apiCfg, w, r)
	})
	v1Router.Get("/feeds", func(w http.ResponseWriter, r *http.Request) {
		handler.HandlerGetFeed(apiCfg, w, r)
	})

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + serverPort,
	}

	log.Printf("Server starting on PORT %s", serverPort)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server listening on PORT %s", serverPort)

}
