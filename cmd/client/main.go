package main

import (
	"context"
	"os/signal"
	"syscall"

	client "github.com/VladKvetkin/gophkeeper/internal/client/app"
	"github.com/VladKvetkin/gophkeeper/internal/client/config"
	"github.com/VladKvetkin/gophkeeper/internal/logger"
)

const (
	CONFIG_PATH = "./config/config.json"
)

func main() {
	if err := start(); err != nil {
		panic(err)
	}
}

func start() error {
	config, err := config.NewConfig(CONFIG_PATH)
	if err != nil {
		return err
	}

	logger, err := logger.NewLogger(config.LogLevel)
	if err != nil {
		return err
	}

	client, err := client.NewClient(logger, config)
	if err != nil {
		return err
	}

	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	defer cancelCtx()

	return client.Start(ctx)
}
