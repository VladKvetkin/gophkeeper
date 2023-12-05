package client

import (
	"fmt"

	"github.com/VladKvetkin/gophkeeper/internal/client/models"
)

func (c *Client) UserAuth() error {
	answer := ""
	fmt.Println("Do you have an account? (y/n)")

	_, err := fmt.Scan(&answer)
	if err != nil {
		return fmt.Errorf("error %w", err)
	}

	switch answer {
	case "y":
		authM, err := buildAuthData(c.printer)
		if err != nil {
			return err
		}

		return c.userSignIn(*authM)

	case "n":
		authM, err := buildAuthData(c.printer)
		if err != nil {
			return err
		}
		return c.userSignUp(*authM)

	default:
		return c.UserAuth()
	}
}

func (c *Client) userSignIn(authData models.AuthModel) error {
	return nil
}

func (c *Client) userSignUp(authData models.AuthModel) error {
	return nil
}

func buildAuthData(p printer) (*models.AuthModel, error) {
	var (
		login, password string
		err             error
	)

	fmt.Println("Please enter your login and password:")

	fmt.Println("login: ")
	_, err = fmt.Scan(&login)
	if err != nil {
		return nil, fmt.Errorf("error %w", err)
	}

	fmt.Println("password: ")
	_, err = fmt.Scan(&password)
	if err != nil {
		return nil, fmt.Errorf("error %w", err)
	}

	return &models.AuthModel{
		Login:    login,
		Password: password,
	}, nil
}
