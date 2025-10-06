package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	client "go_tutorial/http-server/client"
	server "go_tutorial/http-server/server"
)

func main() {

	serverHandler := server.NewServer()
	server := &http.Server{
		Addr:        ":8080",
		Handler:     serverHandler,
		IdleTimeout: 60 * time.Second,
		ReadTimeout: 10 * time.Second,
	}

	go func() {
		fmt.Println("started local server at port 8080")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("failed to initalize a local server", err)
		}
	}()

	time.Sleep(2 * time.Second)

	client.NewClient()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("\nShutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exited gracefully")

}
