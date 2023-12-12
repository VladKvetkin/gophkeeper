package client

import (
	"fmt"
)

// GetUserDataList – получение всех сохранённых данных (мета-данных) пользователя.
func (c *Client) GetUserDataList() error {
	records, err := c.gRPCClient.GetUserDataList()
	if err != nil {
		return err
	}

	if len(records) == 0 {
		return ErrNoData
	}

	fmt.Println("You have these saved data:")
	for _, el := range records {
		fmt.Printf("id: %d, name: %s, type: %s, version: %d\n", el.ID, el.Name, el.DataType, el.Version)
	}

	return nil
}
