package grpcclient

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/VladKvetkin/gophkeeper/internal/client/config"
	"github.com/VladKvetkin/gophkeeper/internal/client/interceptors"

	pb "github.com/VladKvetkin/gophkeeper/pkg/grpc/gophkeeper"
)

// Client – Объект gRPC клиента для общения с сервером.
type Client struct {
	gRPCClient pb.GophKeeperClient
	config     *config.Config
	authToken  string
	timeout    time.Duration
}

// NewGRPCClient – функция инициализации gRPC клиента.
func NewGRPCClient(c *config.Config) (*Client, error) {
	client := &Client{
		config:  c,
		timeout: time.Duration(10) * time.Second,
	}

	conn, err := grpc.Dial(
		c.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.AuthInterceptor(client)),
	)

	if err != nil {
		return nil, fmt.Errorf("gRPC connection refused: %w", err)
	}

	gRPCClient := pb.NewGophKeeperClient(conn)
	client.gRPCClient = gRPCClient

	return client, nil
}

// GetAuthToken – метод получения AuthToken пользователя.
func (c *Client) GetAuthToken() string {
	return c.authToken
}
