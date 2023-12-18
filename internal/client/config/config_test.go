package config

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name     string
		fileData string
		fileErr  bool
		wantErr  bool
	}{
		{
			name:     "success",
			fileErr:  false,
			fileData: `{"address": "localhost:3333"}`,
			wantErr:  false,
		},
		{
			name:    "file open error",
			fileErr: true,
			wantErr: true,
		},
		{
			name:     "parse json error",
			fileErr:  false,
			fileData: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := ""
			if !tt.fileErr {
				dir, err := os.MkdirTemp("", "test_config")
				if err != nil {
					t.Errorf("create temp dir error: %v", err)
				}

				defer func() {
					_ = os.RemoveAll(dir)
				}()

				file, err := os.CreateTemp(dir, "test_config.json")
				if err != nil {
					t.Errorf("create temp file error: %v", err)
				}

				_, _ = file.Write([]byte(tt.fileData))
				path = file.Name()
			}

			_, err := NewConfig(path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
