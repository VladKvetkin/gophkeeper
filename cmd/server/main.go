package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/VladKvetkin/gophkeeper/internal/logger"
	"github.com/VladKvetkin/gophkeeper/internal/server/config"
	"github.com/VladKvetkin/gophkeeper/internal/server/handlers"
	"github.com/VladKvetkin/gophkeeper/internal/server/storage"
)

const (
	CONFIG_PATH = "./config/config.json"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	settings, err := config.NewConfig(CONFIG_PATH)
	if err != nil {
		return err
	}
	l, err := logger.NewLogger(settings.LogLevel)
	if err != nil {
		return err
	}

	rep, err := storage.NewStorage(settings.DatabaseDSN, l)
	if err != nil {
		l.Log.Errorf("database error: %v", err)
		return err
	}
	defer func() {
		if err = rep.Close(); err != nil {
			l.Log.Errorf("database close error: %v", err)
		}
	}()

	server := handlers.NewServer(rep, settings, l)

	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	defer cancelCtx()

	return server.Start(ctx)
}
