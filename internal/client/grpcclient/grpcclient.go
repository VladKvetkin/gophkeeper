package grpcclient

import (
	"fmt"
	"time"

	"github.com/VladKvetkin/gophkeeper/internal/client/config"
	pb "github.com/VladKvetkin/gophkeeper/pkg/grpc/gophkeeper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCClient – Объект gRPC клиента для общения с сервером
type GRPCClient struct {
	gRPCClient pb.GophKeeperClient
	config     *config.Config
	authToken  string
	timeout    time.Duration
}

// NewGRPCClient – конструктор GRPC клиента
func NewGRPCClient(c *config.Config) (*GRPCClient, error) {
	client := &GRPCClient{
		config:  c,
		timeout: time.Duration(10) * time.Second,
	}

	conn, err := grpc.Dial(
		c.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, fmt.Errorf("gRPC connection refused: %w", err)
	}

	gRPCClient := pb.NewGophKeeperClient(conn)
	client.gRPCClient = gRPCClient

	return client, nil
}

// GetAuthToken – метод получения AuthToken пользователя.
func (c *GRPCClient) GetAuthToken() string {
	return c.authToken
}
