package server

import (
	"consoleapp/internal/middleware"
	"net/http"
	"time"
)

var appConfig *AppConfig

func NewServer(cfg *AppConfig) *http.Server {
	appConfig = cfg

	chain := middleware.MiddlewareChain(
		middleware.CORSMiddleware,
		middleware.RequestLoggerMiddleware,
	)

	return &http.Server{
		Addr:         appConfig.getAddress(),
		Handler:      chain(appConfig.RegisterRoutes()),
		ReadTimeout:  time.Duration(appConfig.ReadTimeout) * time.Minute,
		WriteTimeout: time.Duration(appConfig.WriteTimeout) * time.Minute,
		IdleTimeout: time.Minute * 30,
	}
}
