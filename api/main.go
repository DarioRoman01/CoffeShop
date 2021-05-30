package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Haizza1/api/handlers"
)

func main() {
	l := log.New(os.Stdout, "api ", log.LstdFlags)
	port := os.Getenv("PORT")

	ph := handlers.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/products", ph)

	s := http.Server{
		Addr:         port,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Printf("Starting server at port %s\n", port)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting the server: %v", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	log.Println("Got signal: ", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
