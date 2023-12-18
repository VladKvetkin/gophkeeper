package client

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/mailru/easyjson"

	"github.com/VladKvetkin/gophkeeper/internal/client/models"
)

// SaveData – сохранение данных пользователя.
// Сценарий:
//  1. Приложение предлагает пользователю выбрать тип данных, которые надо сохранить.
//  2. Приложение просит пользователя ввести необходимые данные для сохранения.
//  3. Приложение преобразует данные в байты и отправляет на сервер.
func (c *Client) SaveData() error {
	var (
		dataType int
		name     string
	)

	fmt.Println("What data type do you want to save?")

	for i, dt := range [4]string{"password", "card", "file", "text"} {
		fmt.Printf("%v. %v\n", i+1, dt)
	}

	_, err := fmt.Scan(&dataType)
	if err != nil {
		return err
	}

	model, err := buildData(dataType)
	if err != nil {
		return err
	}

	fmt.Println("What name to save the data with?")

	_, err = fmt.Scan(&name)
	if err != nil {
		return err
	}

	model.Name = name

	err = c.gRPCClient.SaveUserData(model)
	if err != nil {
		return err
	}

	fmt.Println("Saved!")

	return nil
}

func buildData(dataType int) (*models.UserData, error) {
	switch dataType {
	case passwordData:
		return buildPassword()
	case cardData:
		return buildCardData()
	case fileData:
		return buildFileData()
	case textData:
		return buildTextData()
	default:
		return nil, errors.New("unknown data type")
	}
}

func buildPassword() (*models.UserData, error) {
	fmt.Println("Please enter password data")

	password := &models.PasswordData{}

	fmt.Println("Site:")
	_, err := fmt.Scan(&password.Site)
	if err != nil {
		return nil, err
	}

	fmt.Println("Login:")
	_, err = fmt.Scanln(&password.Login)
	if err != nil {
		return nil, err
	}

	fmt.Println("Password:")
	_, err = fmt.Scanln(&password.Password)
	if err != nil {
		return nil, err
	}

	byteData, err := easyjson.Marshal(password)
	if err != nil {
		return nil, err
	}

	return &models.UserData{
		DataType: passwordData,
		Data:     byteData,
	}, nil
}

func buildCardData() (*models.UserData, error) {
	fmt.Println("Please enter card data")
	card := &models.CardData{}

	fmt.Println("Card number:")
	_, err := fmt.Scanln(&card.Number)
	if err != nil {
		return nil, err
	}

	fmt.Println("Card exp date:")
	_, err = fmt.Scanln(&card.ExpDate)
	if err != nil {
		return nil, err
	}

	fmt.Println("Card holder:")
	_, err = fmt.Scanln(&card.CardHolder)
	if err != nil {
		return nil, err
	}

	byteData, err := easyjson.Marshal(card)
	if err != nil {
		return nil, err
	}

	return &models.UserData{
		DataType: cardData,
		Data:     byteData,
	}, nil
}

func buildFileData() (*models.UserData, error) {
	fmt.Println("Please enter path to file")
	file := &models.FileData{}

	_, err := fmt.Scanln(&file.Path)
	if err != nil {
		return nil, err
	}

	openedFile, err := os.Open(file.Path)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = openedFile.Close()
	}()

	stat, err := openedFile.Stat()
	if err != nil {
		return nil, err
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(openedFile).Read(bs)
	if err != nil && errors.Is(err, io.EOF) {
		return nil, err
	}

	file.Data = bs

	byteData, err := easyjson.Marshal(file)
	if err != nil {
		return nil, err
	}

	return &models.UserData{
		DataType: fileData,
		Data:     byteData,
	}, nil
}

func buildTextData() (*models.UserData, error) {
	text := &models.TextData{}

	_, err := fmt.Scanln(&text.Text)
	if err != nil {
		return nil, err
	}

	byteData, err := easyjson.Marshal(text)
	if err != nil {
		return nil, err
	}

	return &models.UserData{
		DataType: textData,
		Data:     byteData,
	}, nil
}
