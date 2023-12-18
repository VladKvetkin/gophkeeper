package handlers

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"

	"github.com/VladKvetkin/gophkeeper/internal/logger"
	"github.com/VladKvetkin/gophkeeper/internal/server/config"
	"github.com/VladKvetkin/gophkeeper/internal/server/interceptors"
	"github.com/VladKvetkin/gophkeeper/internal/server/storage"
	crypto "github.com/VladKvetkin/gophkeeper/internal/services/mycrypto"
	pb "github.com/VladKvetkin/gophkeeper/pkg/grpc/gophkeeper"
)

const (
	missingKeyErrText = "missing key in context"
)

type repository interface {
	CreateUser(ctx context.Context, login, password string) (int64, error)
	FindUserByLogin(ctx context.Context, login string) (*storage.User, error)
	SaveUserData(ctx context.Context, userID int64, name string, dataType int64, data []byte) error
	GetUserData(ctx context.Context, userID int64) ([]storage.ShortRecord, error)
	FindUserRecord(ctx context.Context, id, userID int64) (*storage.Record, error)
}

type crypt interface {
	HashFunc(src string) (string, error)
	CompareHash(src, hash string) error
	BuildJWT(userID int64, secret string) (string, error)
	GetUserID(tokenString, secret string) (int64, error)
}

// Server – сервер приложения, который отвечает за хранение и обработку приватных данных пользователя.
type Server struct {
	pb.UnimplementedGophKeeperServer
	Storage repository
	crypto  crypt
	Config  *config.Config
	Logger  *logger.Logger
}

// NewServer – функция инициализации сервера.
// Функция принимает репозиторий, конфигуратор и логгер.
func NewServer(r repository, c *config.Config, l *logger.Logger) *Server {
	return &Server{
		Storage: r,
		Config:  c,
		Logger:  l,
		crypto:  &crypto.MyCrypt{},
	}
}

// Start – метод для запуска сервера приложения.
func (s *Server) Start(ctx context.Context) error {
	listen, err := net.Listen("tcp", s.Config.Address)
	if err != nil {
		s.Logger.Log.Error(err)
		return fmt.Errorf("tcp connection failed")
	}

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptors.AuthInterceptor(s.Logger, s.Config.SecretKey, s.crypto),
		logging.UnaryServerInterceptor(interceptors.LoggerInterceptor()),
	))

	pb.RegisterGophKeeperServer(gRPCServer, s)

	wg := &sync.WaitGroup{}
	defer func() {
		wg.Wait()
	}()

	wg.Add(1)
	go func() {
		defer s.Logger.Log.Info("server has been shutdown")
		defer wg.Done()
		<-ctx.Done()

		s.Logger.Log.Info("app got a signal")
		gRPCServer.GracefulStop()
	}()

	s.Logger.Log.Info("gRPC server is running")

	return gRPCServer.Serve(listen)
}
