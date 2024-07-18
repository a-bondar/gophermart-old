package main

import (
	"log"
	"net/http"

	"github.com/a-bondar/gophermart/internal/config"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	cfg := config.NewConfig()
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://example.org", 307)

	mux.Handle("/foo", rh)

	log.Printf("Starting serve on %s", cfg.RunAddr)

	return http.ListenAndServe(cfg.RunAddr, mux)
}
