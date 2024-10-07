package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/prasad89/go-restful-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go server!"))
	})

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("Server is started")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}
}
