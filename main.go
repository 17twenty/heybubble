package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

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

		type Message struct {
			HTMLContent string
			Author      string
			Timestamp   time.Time
			IsAuthor    bool
		}

		msg := []string{
			`<p>Hey 👋 - It's Nick, Welcome to the site.</p>`,
			`<p>I'm currently the Chief Technology and Product Officer at a company called FeeWise.</p>`,
			`<p>You can find me at <a class="text-blue-600 text-blue-500 hover:underline" href="https://www.curiola.com">https://www.curiola.com</a> or you can drop me an email at <a class="text-blue-600 dark:text-blue-500 hover:underline" href="mailto:nick@curiola.com">nick@curiola.com</a>.</p>`,
			`<p>I have worked across a number of technology and finance companies to build and execute on products and designs and currently act as an advisor and fractional CTO for a number of other startups here in Sydney, Australia.</p>`,
			`<p>So if you want to chat, reach out or check me out on LinkedIn at <a class="text-blue-600 hover:underline" href="https://www.linkedin.com/in/nickglynn/">https://www.linkedin.com/in/nickglynn/</a>.</p>`,
		}

		// If there's more, return thinking...
		if offset < len(msg) {
			if offset >= 0 {
				partialEncoder(w, "bubbleleft", msg[offset])
				// partialEncoder(w, "bubbleright", "Cool!")
			}
			partialEncoder(w, "thinking", struct {
				Offset int
			}{
				Offset: offset + 1,
			})
		} else {
			// This is the default to tell htmx to stop polling
			w.WriteHeader(286)
		}
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, r))
}
