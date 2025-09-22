package config

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Server represents HTTP server configuration
type Server struct {
	httpServer *http.Server
	config     *Config
}

// NewServer creates a new server instance
func NewServer(engine *gin.Engine, config *Config) *Server {
	srv := &http.Server{
		Addr:         ":" + config.Server.Port,
		Handler:      engine,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{
		httpServer: srv,
		config:     config,
	}
}

// Start starts the HTTP server with graceful shutdown
func (s *Server) Start() error {
	// Start server in a goroutine
	go func() {
		log.Printf("Ghoona Camp Backend server starting on port %s", s.config.Server.Port)
		log.Printf("Environment: %s", s.config.Server.Environment)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Give outstanding requests a deadline for completion
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	log.Println("Server exited")
	return nil
}

// Stop stops the HTTP server
func (s *Server) Stop() error {
	return s.httpServer.Close()
}