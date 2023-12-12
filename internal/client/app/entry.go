package client

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

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

			commandNumber := 0
			_, err := fmt.Scan(&commandNumber)
			if err != nil {
				return err
			}

			switch commandNumber {
			case getUserDataList:
				err := c.GetUserDataList()
				if err != nil {
					if errors.Is(err, ErrNoData) {
						fmt.Println("You have no saved data")
						continue
					}
					c.Logger.Log.Error(err)
					continue
				}
			case getUserData:
				err := c.GetUserData()
				if err != nil {
					c.Logger.Log.Error(err)
					continue
				}
			case saveUserData:
				err := c.SaveData()
				if err != nil {
					c.Logger.Log.Error(err)
					continue
				}
			default:
				fmt.Println("Unknown command")
			}

			fmt.Printf("\n====================\n\n")
		}
	}
}
