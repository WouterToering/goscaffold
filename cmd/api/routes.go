package api

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func NewHandler(s *Server, logger *zap.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Use(requestLogger(logger))
	r.Use(middleware.Recoverer)

	r.Get("/getPotato", s.GetPotato())

	logRoutes(r, logger)

	return r
}

func logRoutes(r *chi.Mux, logger *zap.Logger) {
	logger.Info("Printing routes")
	if err := chi.Walk(r, printRoute(logger)); err != nil {
		logger.Error("Failed to walk routes:", zap.Error(err))
	}
}

func printRoute(logger *zap.Logger) chi.WalkFunc {
	return func(
		method string,
		route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler,
	) error {
		route = strings.Replace(route, "/*/", "/", -1)
		logger.Info("Registering route", zap.String("method", method), zap.String("route", route))
		return nil
	}
}
