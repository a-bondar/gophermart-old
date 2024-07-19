package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/a-bondar/gophermart/internal/config"
	"github.com/jackc/pgx/v5"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	cfg := config.NewConfig()
	mux := http.NewServeMux()
	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURI)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("unable to ping database: %w", err)
	}

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		err := conn.Ping(r.Context())
		if err != nil {
			http.Error(w, "database is not available", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Printf("Starting server on %s", cfg.RunAddr)

	return http.ListenAndServe(cfg.RunAddr, mux)
}
