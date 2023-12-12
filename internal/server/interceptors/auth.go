package interceptors

import (
	"context"
	"net/http"

	"github.com/VladKvetkin/gophkeeper/internal/logger"
	"github.com/VladKvetkin/gophkeeper/internal/services"
	"github.com/VladKvetkin/gophkeeper/pkg/grpc/gophkeeper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type crypt interface {
	GetUserID(string, string) (int64, error)
}

// AuthInterceptor – интерсептор сервера для проверки авторизации пользователя.
func AuthInterceptor(l *logger.Logger, secret string, cr crypt) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, r interface{}, i *grpc.UnaryServerInfo, h grpc.UnaryHandler) (interface{}, error) {
		if i.FullMethod == gophkeeper.GophKeeper_SignUp_FullMethodName ||
			i.FullMethod == gophkeeper.GophKeeper_SignIn_FullMethodName {
			return h(ctx, r)
		}

		var token string
		if meta, ok := metadata.FromIncomingContext(ctx); ok {
			values := meta.Get("token")
			if len(values) > 0 {
				token = values[0]
			}
		}
		if len(token) == 0 {
			return nil, status.Error(codes.Unauthenticated, http.StatusText(http.StatusForbidden))
		}

		userID, err := cr.GetUserID(token, secret)
		if err != nil {
			l.Log.Debugf("invalid token: %v", token)
			return nil, status.Error(codes.Unauthenticated, http.StatusText(http.StatusForbidden))
		}

		ctx = context.WithValue(ctx, services.UserIDContextKey, userID)

		return h(ctx, r)
	}
}
