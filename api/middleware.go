package api

import (
	"log"
	"net/http"
	"time"
)

func WithRequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
func JSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func (s *Server) Routes() http.Handler                   { return WithRequestLog(JSON(s.Handler())) }
func status(w http.ResponseWriter, code int, msg string) { w.WriteHeader(code); w.Write([]byte(msg)) }
