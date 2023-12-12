package client

import (
	"github.com/VladKvetkin/gophkeeper/internal/client/config"
	"github.com/VladKvetkin/gophkeeper/internal/client/grpcclient"
	"github.com/VladKvetkin/gophkeeper/internal/client/models"
	"github.com/VladKvetkin/gophkeeper/internal/logger"
)

type grpcClient interface {
	SignIn(model models.AuthModel) (models.AuthToken, error)
	SignUp(model models.AuthModel) (models.AuthToken, error)
	GetUserData(model models.UserDataModel) (*models.UserData, error)
	GetUserDataList() ([]models.UserDataList, error)
	SaveUserData(model *models.UserData) error
}

type clientCache interface {
	Append(data *models.UserData)
	GetUserData(model models.UserDataModel) (*models.UserData, error)
	GetUserDataList() []models.UserDataList
}

// Client – структура, которая служит для работы с сервером. Отвечает за сценарий приложения.
type Client struct {
	gRPCClient   grpcClient
	Config       *config.Config
	Logger       *logger.Logger
	dataSyncChan chan int64
}

// NewClient – функция для инициализации клиента.
// Принимает логгер, для внутреннего логгирования и конфигурацию для клиента.
func NewClient(l *logger.Logger, c *config.Config) (*Client, error) {
	gRPCClient, err := grpcclient.NewGRPCClient(c)
	if err != nil {
		return nil, err
	}

	client := &Client{
		gRPCClient: gRPCClient,
		Config:     c,
		Logger:     l,
	}

	return client, nil
}
