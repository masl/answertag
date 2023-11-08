package main

import (
	"embed"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/masl/answertag/storage/inmemory"
	"github.com/masl/answertag/tmpl"
	"github.com/masl/answertag/web"
	"github.com/masl/answertag/ws"
)

var (
	//go:embed templates/*
	templateFS embed.FS

	//go:embed static/*
	staticFS embed.FS
)

func main() {
	handleSignal()

	tm := tmpl.NewTemplateManager()
	err := tmpl.RegisterTemplates(tm, templateFS)
	if err != nil {
		slog.Error("error registering templates", "error", err)
		os.Exit(1)
	}

	store := inmemory.New()
	hub := ws.NewHub(store, tm)
	go hub.Run()

	server := &http.Server{
		Addr:    ":3000",
		Handler: web.GetRouter(store, tm, staticFS, hub),
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
