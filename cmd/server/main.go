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

	"github.com/Listantiyo/pos-system/internal/config"
	"github.com/Listantiyo/pos-system/internal/database"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Connect to database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 3. Test Redis Connection
	ctx := context.Background()
	if err := db.Redis.Ping(ctx).Err; err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	log.Println("✅ Connected to Redis")

	// 4. Setup graceful shutdown
	adr := fmt.Sprintf(":%s", cfg.AppPort)
	srv := &http.Server{
		Addr:	adr,
		Handler: nil,
	}

	// Channel untuk graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server di goroutine
	go func() {
		log.Printf("🚀 Server starting on http://localhost:%s", cfg.AppPort)
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("Server error", err)
		}
	}()

	<-stop
	log.Println("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown error:", err)
	}

	log.Println("Server stop gracefully")
}