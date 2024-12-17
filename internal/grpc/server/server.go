package server

import (
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/rodeorm/keeper/internal/cfg"
	"github.com/rodeorm/keeper/internal/grpc/interc"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"github.com/rodeorm/keeper/internal/logger"
)

type grpcServer struct {
	srv *grpc.Server // общая реализация grpc сервера
	cfg *cfg.Server  // часть сервера для приложения
	proto.UnimplementedKeeperServiceServer
}

// Start запускает grpc-сервер
func Start(cfg *cfg.Server, wg *sync.WaitGroup, exit chan struct{}) error {
	grpcSrv := grpcServer{cfg: cfg}
	listen, err := net.Listen("tcp", cfg.RunAddress)

	if err != nil {
		logger.Error("ServerStart", "ошибка при попытке начать слушать порт", err.Error())
		return err
	}

	grpcSrv.srv = grpc.NewServer(grpc.UnaryInterceptor(interc.UnaryServerLogInterceptor))

	proto.RegisterKeeperServiceServer(grpcSrv.srv, &grpcSrv)

	logger.Log.Info("grpc server started",
		zap.String("server gRPC has started on the port: ", cfg.ServerConfig.RunAddress),
	)

	var g errgroup.Group

	g.Go(func() error {

		if err := grpcSrv.srv.Serve(listen); err != nil {
			return err
		}
		return nil
	})

	// Обработка сигналов
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	logger.Log.Info("grpc server gracefuly shutdown",
		zap.String("grpc server начал изящно завершать работу", cfg.ServerConfig.RunAddress),
	)
	grpcSrv.srv.GracefulStop() // Корректное завершение работы сервера

	logger.Log.Info("grpc server has gracefuly shutdowned",
		zap.String("grpc server изящно завершил работу", cfg.ServerConfig.RunAddress),
	)
	defer close(exit)
	defer wg.Done()

	// Ожидание завершения всех горутин
	if err := g.Wait(); err != nil {
		logger.Error("ServerStart", "ошибка при запуске grpc сервера", err.Error())
		return err
	}

	return nil
}
