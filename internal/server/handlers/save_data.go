package handlers

import (
	"context"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/VladKvetkin/gophkeeper/internal/services"
	pb "github.com/VladKvetkin/gophkeeper/pkg/grpc/gophkeeper"
)

// SaveData – метод сохранения данных пользователя на сервере.
func (s *Server) SaveData(ctx context.Context, in *pb.SaveDataRequest) (*pb.SaveDataResponse, error) {
	userID, ok := ctx.Value(services.UserIDContextKey).(int64)
	if !ok {
		s.Logger.Log.Error(missingKeyErrText)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	if err := s.Storage.SaveUserData(ctx, userID, in.Name, in.DataType, in.Data); err != nil {
		s.Logger.Log.Error(err)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	return &pb.SaveDataResponse{
		Result: "OK",
	}, nil
}
