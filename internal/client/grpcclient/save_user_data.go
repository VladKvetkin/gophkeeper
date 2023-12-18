package grpcclient

import (
	"context"
	"fmt"

	"github.com/VladKvetkin/gophkeeper/internal/client/models"
	pb "github.com/VladKvetkin/gophkeeper/pkg/grpc/gophkeeper"
)

// SaveUserData – метод сохранения данных пользователя на сервер.
func (c *Client) SaveUserData(model *models.UserData) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	req := &pb.SaveDataRequest{
		Name:     model.Name,
		Data:     model.Data,
		DataType: model.DataType,
	}
	_, err := c.gRPCClient.SaveData(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC SaveUserData error: %w", err)
	}

	return nil
}
