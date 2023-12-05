package client

import (
	"context"
	"fmt"

	"github.com/VladKvetkin/gophkeeper/internal/client/config"
	"github.com/VladKvetkin/gophkeeper/internal/client/grpcclient"
	"github.com/VladKvetkin/gophkeeper/internal/logger"
	"golang.org/x/sync/errgroup"
)

type printer interface {
	Print(s string)
	Scan(a ...interface{}) (int, error)
}

// Client – структура клиента
type Client struct {
	gRPCClient   *grpcclient.GRPCClient
	printer      printer
	Config       *config.Config
	Logger       *logger.Logger
	dataSyncChan chan int64
}

// NewClient – конструктор клиента
// Принимает логгер, для логгирования и конфигурацию.
func NewClient(logger *logger.Logger, config *config.Config) (*Client, error) {
	gRPCClient, err := grpcclient.NewGRPCClient(config)
	if err != nil {
		return nil, err
	}

	client := &Client{
		gRPCClient: gRPCClient,
		Config:     config,
		Logger:     logger,
	}

	return client, nil
}

// Start – функция для начала работы с клиентом.
// Сначала пользователь должен авторизоваться.
// Затем пользователю предлагают выбрать одну из команд:
//  1. Получение всех сохранённых текстовых данных.
//  2. Получение бинарных данных.
//  3. Сохранение данных.
//  4. Редактирование данных.
func (c *Client) Start(ctx context.Context) error {
	fmt.Println("Hello! I'm GophKeeper. I can save your private information.")

	if err := c.UserAuth(); err != nil {
		c.Logger.Log.Error(err)
		return err
	}

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return c.startSession(ctx)
	})

	return eg.Wait()
}

func (c *Client) startSession(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			c.Logger.Log.Info("client has been shutdown")
			return nil
		default:
			fmt.Println("Choose command (enter number of command)")
			fmt.Println("1. Get all data")
			fmt.Println("2. Get binary data")
			fmt.Println("3. Save some data")
			fmt.Println("4. Edit saved data")

			commandNumber := 0
			_, err := fmt.Scan(&commandNumber)
			if err != nil {
				return err
			}

			switch commandNumber {

			default:
				fmt.Println("Unknown command")
			}

			fmt.Printf("\n====================\n\n")
		}
	}
}
