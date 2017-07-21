package main

import (
	"context"
	"net/http"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func healthcheck_handler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
}

func default_handler(res http.ResponseWriter, req *http.Request) {
	log.Printf("not found: %s", req.URL.RequestURI())
	res.WriteHeader(http.StatusNotFound)
}

func main() {
	http.HandleFunc("/status", healthcheck_handler)
	http.HandleFunc("/", default_handler)

	srv := &http.Server{
		Handler: nil,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println()
		log.Printf("Received signal: %v", sig)
		log.Println("Shutting down...")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		srv.Shutdown(ctx)
	}()

	log.Println("Listening on addr: :8080")
	log.Fatal(srv.ListenAndServe())
}
