package client

import (
	"fmt"

	"github.com/VladKvetkin/gophkeeper/internal/client/models"
)

// UserAuth – функция авторизации пользователя.
func (c *Client) UserAuth() error {
	var ans string
	fmt.Println("Do you have an account? (y/n)")

	_, err := fmt.Scan(&ans)
	if err != nil {
		return err
	}

	switch ans {
	case "y":
		authM, err := buildAuthData()
		if err != nil {
			return err
		}
		return c.userSignIn(*authM)
	case "n":
		authM, err := buildAuthData()
		if err != nil {
			return err
		}
		return c.userSignUp(*authM)
	default:
		return c.UserAuth()
	}
}

func (c *Client) userSignIn(authM models.AuthModel) error {
	_, err := c.gRPCClient.SignIn(authM)
	if err != nil {
		return fmt.Errorf("SignIn error: %w", err)
	}

	return nil
}

func (c *Client) userSignUp(authM models.AuthModel) error {
	_, err := c.gRPCClient.SignUp(authM)
	if err != nil {
		return fmt.Errorf("SignUp error: %w", err)
	}

	return nil
}

func buildAuthData() (*models.AuthModel, error) {
	var (
		login, password string
		err             error
	)

	fmt.Println("Please enter your login and password.")
	fmt.Println("Login: ")
	_, err = fmt.Scan(&login)
	if err != nil {
		return nil, err
	}

	fmt.Println("Password: ")
	_, err = fmt.Scan(&password)
	if err != nil {
		return nil, err
	}

	return &models.AuthModel{
		Login:    login,
		Password: password,
	}, nil
}
