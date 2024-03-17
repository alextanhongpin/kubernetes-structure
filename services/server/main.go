package main

import (
	"fmt"
	"net/http"
	"os"

	"log/slog"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request received", slog.String("method", r.Method), slog.String("url", r.URL.String()))
		fmt.Fprintf(w, "Hello, World! - %s", env)
	}))
	logger.Info("Listening on :8080")
	panic(http.ListenAndServe(":8080", mux))
}
