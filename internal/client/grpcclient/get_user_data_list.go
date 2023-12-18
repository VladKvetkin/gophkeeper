package grpcclient

import (
	"context"
	"fmt"

	"github.com/VladKvetkin/gophkeeper/internal/client/models"
	pb "github.com/VladKvetkin/gophkeeper/pkg/grpc/gophkeeper"
)

// GetUserDataList – метод получения всех сохранённых данных (мета-данных) пользователя с сервера.
func (c *Client) GetUserDataList() ([]models.UserDataList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	req := &pb.UserDataListRequest{}
	res, err := c.gRPCClient.GetUserDataList(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC GetUserDataList error: %w", err)
	}

	records := make([]models.UserDataList, 0, len(res.Data))
	for _, el := range res.Data {
		rec := models.UserDataList{
			ID:       el.Id,
			Name:     el.Name,
			DataType: el.DataType,
			Version:  el.Version,
		}
		records = append(records, rec)
	}

	return records, nil
}
