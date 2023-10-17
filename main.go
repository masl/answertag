package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handleSignal()

	slog.Info("Hello, World!")
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
