package grpc

import (
	"book/book_go_book_service/config"
	"book/book_go_book_service/genproto/book_service"
	"book/book_go_book_service/grpc/client"
	"book/book_go_book_service/grpc/service"
	"book/book_go_book_service/pkg/logger"
	"book/book_go_book_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	book_service.RegisterBookServiceServer(grpcServer, service.NewBookService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
