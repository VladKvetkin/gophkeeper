package client

import (
	"errors"
	"net"
	"os"
	"testing"

	"github.com/mailru/easyjson"

	"github.com/VladKvetkin/gophkeeper/internal/client/config"
	"github.com/VladKvetkin/gophkeeper/internal/client/models"
	"github.com/VladKvetkin/gophkeeper/internal/logger"
)

var (
	testLogger *logger.Logger
	testConfig *config.Config
	testClient *Client
	testData   *models.UserData
)

func TestMain(m *testing.M) {
	var err error
	testLogger, err = logger.NewLogger("info")
	if err != nil {
		panic(err)
	}

	testConfig = &config.Config{
		Address: "localhost:3333",
	}

	testClient = &Client{
		Config:       testConfig,
		Logger:       testLogger,
		gRPCClient:   &testGRPCClient{},
		dataSyncChan: make(chan int64),
	}

	passData, _ := easyjson.Marshal(&models.PasswordData{
		Site:     "test.com",
		Login:    "testLogin",
		Password: "testPass",
	})
	testData = &models.UserData{
		Name:     "testName",
		DataType: 1,
		Data:     passData,
		ID:       1,
		Version:  1,
	}

	code := m.Run()

	os.Exit(code)
}

func TestNewClient(t *testing.T) {
	type args struct {
		l *logger.Logger
		c *config.Config
	}
	tests := []struct {
		name    string
		args    args
		grpcRun bool
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				l: &logger.Logger{},
				c: &config.Config{
					Address: "localhost:3333",
				},
			},
			grpcRun: true,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				l: &logger.Logger{},
				c: &config.Config{},
			},
			grpcRun: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.grpcRun {
				conn, err := net.Listen("tcp", ":3200")
				if err != nil {
					t.Errorf("gRPC server start error = %v", err)
				}
				defer func() {
					_ = conn.Close()
				}()
			}

			_, err := NewClient(tt.args.l, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

type testGRPCClient struct{}

func (t *testGRPCClient) SignIn(model models.AuthModel) (models.AuthToken, error) {
	if model.Login == "errorLogin" {
		return "", errors.New(`request error`)
	}
	return "testToken", nil
}

func (t *testGRPCClient) SignUp(model models.AuthModel) (models.AuthToken, error) {
	if model.Login == "errorLogin" {
		return "", errors.New(`request error`)
	}
	return "testToken", nil
}

func (t *testGRPCClient) GetUserData(model models.UserDataModel) (*models.UserData, error) {
	if model.ID == 2 {
		return nil, errors.New(`request error`)
	}
	return testData, nil
}

func (t *testGRPCClient) GetUserDataList() ([]models.UserDataList, error) {
	return []models.UserDataList{{
		Name:     "testName",
		DataType: 1,
		ID:       1,
		Version:  1,
	}}, nil
}

func (t *testGRPCClient) SaveUserData(model *models.UserData) error {
	if model.ID == 2 {
		return errors.New(`request error`)
	}
	return nil
}
