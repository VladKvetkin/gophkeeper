package client

import (
	"testing"

	"github.com/VladKvetkin/gophkeeper/internal/client/config"
	"github.com/VladKvetkin/gophkeeper/internal/logger"
)

func TestClient_SaveData(t *testing.T) {
	type fields struct {
		gRPCClient   grpcClient
		Config       *config.Config
		Logger       *logger.Logger
		dataSyncChan chan int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "error",
			fields: fields{
				gRPCClient: &testGRPCClient{},
				Config:     testConfig,
				Logger:     testLogger,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				gRPCClient:   tt.fields.gRPCClient,
				Config:       tt.fields.Config,
				Logger:       tt.fields.Logger,
				dataSyncChan: tt.fields.dataSyncChan,
			}
			if err := c.SaveData(); (err != nil) != tt.wantErr {
				t.Errorf("SaveData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_buildData(t *testing.T) {
	type args struct {
		dataType int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "build password",
			args: args{
				dataType: 1,
			},
			wantErr: false,
		},
		{
			name: "build card",
			args: args{
				dataType: 2,
			},
			wantErr: false,
		},
		{
			name: "build file",
			args: args{
				dataType: 3,
			},
			wantErr: true,
		},
		{
			name: "build text",
			args: args{
				dataType: 4,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := buildData(tt.args.dataType)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
