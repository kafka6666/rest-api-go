package main

import (
	"fmt"
	"github.com/kafka6666/rest-api-go/internal/config"
	"log"
	"net/http"
)

func main() {

	// load config file
	cfg := config.MustLoad()
	// database setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Welcome to Students API"))
		if err != nil {
			log.Fatal(err)
		}
	})
	// setup server
	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("Server started on %s", cfg.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
