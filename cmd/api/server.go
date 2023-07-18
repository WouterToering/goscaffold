package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"jemoeder.com/internal/services"
)

type Server struct {
	l *zap.Logger
	*services.Services
	*http.Server
}

func NewServer(listenaddr string) (*Server, error) {
	services := services.NewServices()

	logger, err := zap.NewProduction(zap.AddStacktrace(zapcore.FatalLevel))
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("Configure API")

	errorLog, _ := zap.NewStdLogAt(logger, zap.ErrorLevel)
	srv := http.Server{
		Addr:         listenaddr,
		ErrorLog:     errorLog,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	server := &Server{logger, services, &srv}
	srv.Handler = NewHandler(server, logger)

	return server, nil
}

func (srv *Server) Start() {
	srv.l.Info("Starting server...")
	defer srv.l.Sync()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.l.Fatal("Could not listen on", zap.String("addr", srv.Addr), zap.Error(err))
		}
	}()
	srv.l.Info("Server is ready to handle requests", zap.String("addr", srv.Addr))
	srv.gracefulShutdown()
}

func (srv *Server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	srv.l.Info("Server is shutting down", zap.String("reason", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		srv.l.Fatal("Could not gracefuly shutdown the server", zap.Error(err))
	}
	srv.l.Info("Server stopped")
}
