package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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
		tmpl := template.Must(template.ParseFiles("templates/index.tpl.html", "partials/thinking.tpl.html"))
		err := tmpl.Lookup("index.tpl.html").Execute(w, "0")
		if err != nil {
			logger.Error("Failed to execute template", "template", tmpl.Name, "error", err)
		}
	})

	r.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		// Get message
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		// log.Println(offset)
		msg := ""

		switch offset {
		case 1:
			msg = `<p>Hey 👋</p>
<p>It's me Nick - thanks for stopping by!</p>
<img class="object-none rounded-lg" src="data:image/jpeg;base64,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">`
		case 2:
			msg = `<p>I'm currently the Chief Technology and Product Officer at a company called FeeWise.</p>`
		case 3:
			msg = `<p>You can find me at <a 
			class="text-blue-600 dark:text-blue-500 hover:underline"
			href="https://www.curiola.com">https://www.curiola.com</a> or you can drop me an email at <a class="text-blue-600 dark:text-blue-500 hover:underline" href="mailto:nick@curiola.com">nick@curiola.com</a>.</p>`
		}

		// If there's more, return thinking...
		if offset < 4 {
			if offset > 0 {
				partialEncoder(w, "bubbleleft", msg)
			}
			partialEncoder(w, "thinking", struct {
				Offset int
			}{
				Offset: offset + 1,
			})
		} else {
			w.WriteHeader(286)
		}
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, r))
}
