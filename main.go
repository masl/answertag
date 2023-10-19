package main

import (
	"embed"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"

	"github.com/masl/answertag/storage/inmemory"
	"github.com/masl/answertag/web"
)

var (
	//go:embed templates/*
	templateFS embed.FS

	//go:embed static/*
	staticFS embed.FS

	// parsed templates
	htmlTemplates *template.Template
)

func main() {
	handleSignal()

	err := parseTemplates()
	if err != nil {
		panic(err)
	}

	store := inmemory.New()

	server := &http.Server{
		Addr:    ":3000",
		Handler: web.GetRouter(store, htmlTemplates, staticFS),
	}

	slog.Info("web server listening on port 3000")
	panic(server.ListenAndServe())
}

func handleSignal() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		slog.Info("exiting, bye-bye!")
		os.Exit(1)
	}()
}

func parseTemplates() (err error) {
	htmlTemplates, err = template.ParseFS(templateFS, "templates/*.html")
	if err != nil {
		return err
	}

	return nil
}
