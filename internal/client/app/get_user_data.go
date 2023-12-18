package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mailru/easyjson"

	"github.com/VladKvetkin/gophkeeper/internal/client/models"
)

// GetUserData – получение сохранённых данных.
// Сценарий:
//  1. Приложение забирает с сервера все сохранённые данные пользователя (мета данные).
//  2. Приложение предлагает пользователю выбрать из сохранённых данных те, которые пользователь хочет получить.
//  3. Приложение достаёт нужные данные и отдаёт пользователю.
func (c *Client) GetUserData() error {
	err := c.GetUserDataList()
	if err != nil {
		if errors.Is(err, ErrNoData) {
			fmt.Println("You have no saved data")
			return nil
		}
		return err
	}

	fmt.Println("Please enter data id")

	var (
		data   *models.UserData
		dataID int64
	)

	_, err = fmt.Scanln(&dataID)
	if err != nil {
		return err
	}

	m := models.UserDataModel{ID: dataID}
	data, err = c.gRPCClient.GetUserData(m)
	if err != nil {
		return err
	}

	err = printData(data)
	if err != nil {
		return err
	}

	return nil
}

func printData(data *models.UserData) error {
	var pretty []byte

	switch data.DataType {
	case passwordData:
		password := &models.PasswordData{}
		err := easyjson.Unmarshal(data.Data, password)

		if err != nil {
			return err
		}

		pretty, err = json.MarshalIndent(password, "", "  ")
		if err != nil {
			return err
		}
	case cardData:
		card := &models.CardData{}
		err := easyjson.Unmarshal(data.Data, card)
		if err != nil {
			return err
		}

		pretty, err = json.MarshalIndent(card, "", "  ")
		if err != nil {
			return err
		}
	case fileData:
		file := &models.FileData{}
		err := easyjson.Unmarshal(data.Data, file)
		if err != nil {
			return err
		}

		pretty, err = json.MarshalIndent(file, "", "  ")
		if err != nil {
			return err
		}
	case textData:
		text := &models.TextData{}
		err := easyjson.Unmarshal(data.Data, text)
		if err != nil {
			return err
		}

		pretty, err = json.MarshalIndent(text, "", "  ")
		if err != nil {
			return err
		}
	default:
		return nil
	}

	fmt.Printf("\nYour data is:\n%s", pretty)

	return nil
}
