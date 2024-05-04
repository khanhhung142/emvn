package server

import (
	"context"
	"emvn/config"
	"emvn/database/nosql/mongodb"
	"emvn/pkg/logger"
	"emvn/pkg/storage/local"
	"emvn/pkg/validator"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartServer() {
	// Load config
	config.InitConfig()
	cfg := config.GetConfig()

	logger.NewLogger(cfg.Log)
	ctx := context.Background()

	mongodb.InitClient(ctx)
	local.InitLocalStorage()
	validator.InitValidator()

	// Register all dependencies
	Register()

	r := InitHandler()

	srv := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	<-ctx.Done()
	log.Println("Server shutdown gracefully.")
}
