package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
)

var (
	logger *slog.Logger
)

func main() {
	port := "8080"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	appEnv := "dev"

	if fromEnv := os.Getenv("ENV"); fromEnv != "" {
		appEnv = fromEnv
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug, // we should toggle this if we're in prod
	}

	var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)
	if appEnv == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}
	logger = slog.New(handler)

	logger.Info("Starting server...", "server", fmt.Sprintf("http://0.0.0.0:%s", port))

	r := mux.NewRouter()

	// Set no caching
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
			wr.Header().Set("Cache-Control", "max-age=0, must-revalidate")
			next.ServeHTTP(wr, req)
		})
	})

	// Setup filehandling
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))

	// Entry route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host)
		tmpl := template.Must(template.ParseFiles("templates/index.tpl.html"))
		err := tmpl.Lookup("index.tpl.html").Execute(w, nil)
		if err != nil {
			logger.Error("Failed to execute template", "template", tmpl.Name, "error", err)
		}
	})

	r.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		partialEncoder(w, "bubble", nil)
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, r))
}
