package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/go-webauthn/webauthn/webauthn"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	// Configuración HTTPS
	certFile = "cert.pem"
	keyFile  = "key.pem"

	// OAuth2
	oauthConfig = &oauth2.Config{
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		RedirectURL:  "https://localhost:8080/auth/callback",
		Scopes:       []string{"email"},
		Endpoint:     google.Endpoint,
	}

	// WebAuthn
	webAuthn *webauthn.WebAuthn

	// API Keys con niveles
	apiKeys = map[string]int{
		"nivel1": 1,
		"nivel2": 2,
	}

	// Rate limiting
	rateLimiter = NewRateLimiter(100, time.Minute)
)

func main() {
	setupWebAuthn()

	mux := http.NewServeMux()

	// Versión 1 de la API
	mux.HandleFunc("/v1/protected", applyMiddlewares(
		protectedHandler,
		whitelistMiddleware,
		rateLimiter.Middleware,
		oauthMiddleware,
	))

	// WebAuthn endpoints
	mux.HandleFunc("/webauthn/register/begin", beginRegistration)
	mux.HandleFunc("/webauthn/register/finish", finishRegistration)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	log.Fatal(server.ListenAndServeTLS(certFile, keyFile))
}

func setupWebAuthn() {
	wconfig := &webauthn.Config{
		RPDisplayName: "Mi Empresa",
		RPID:          "localhost",
		RPOrigins:     []string{"https://localhost:8080"},
	}
	var err error
	webAuthn, err = webauthn.New(wconfig)
	if err != nil {
		log.Fatal(err)
	}
}

// Middleware para Whitelisting
func whitelistMiddleware(next http.HandlerFunc) http.HandlerFunc {
	whitelist := []string{"127.0.0.1", "192.168.1.0/24"}
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		for _, cidr := range whitelist {
			_, subnet, _ := net.ParseCIDR(cidr)
			if subnet.Contains(net.ParseIP(ip)) {
				next(w, r)
				return
			}
		}
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

// Middleware para API Keys niveladas
func apiKeyMiddleware(requiredLevel int) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")
			if level, ok := apiKeys[apiKey]; !ok || level < requiredLevel {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next(w, r)
		}
	}
}

// Rate Limiter
type RateLimiter struct {
	mu     sync.Mutex
	limits map[string]int
	reset  time.Time
	max    int
	window time.Duration
}

func NewRateLimiter(max int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limits: make(map[string]int),
		max:    max,
		window: window,
		reset:  time.Now().Add(window),
	}
}

func (rl *RateLimiter) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rl.mu.Lock()
		defer rl.mu.Unlock()

		if time.Now().After(rl.reset) {
			rl.limits = make(map[string]int)
			rl.reset = time.Now().Add(rl.window)
		}

		identifier := r.Header.Get("X-API-Key")
		if identifier == "" {
			identifier, _, _ = net.SplitHostPort(r.RemoteAddr)
		}

		if rl.limits[identifier] >= rl.max {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		rl.limits[identifier]++
		next(w, r)
	}
}

// Middleware OAuth2
func oauthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if !validOAuthToken(token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func validOAuthToken(token string) bool {
	// Implementar validación real del token
	return true
}

// Handler de ejemplo
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Acceso autorizado")
}

// WebAuthn Handlers (simplificados)
func beginRegistration(w http.ResponseWriter, r *http.Request) {
	// Implementación real necesaria
}

func finishRegistration(w http.ResponseWriter, r *http.Request) {
	// Implementación real necesaria
}

// Helpers para middlewares
func applyMiddlewares(h http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, mw := range middlewares {
		h = mw(h)
	}
	return h
}
