package cmd

import (
	"context"
	"log"
	"myChat/config"
	"myChat/internal"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg := config.NewConfig()
	time.Local = time.UTC

	// Cancel the context on system signals (graceful shutdown).
	// SIGTERM is sent by k8s to terminate the process.
	rootCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()

	// Create modules
	internal.InitModules(r, rootCtx, cfg)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
