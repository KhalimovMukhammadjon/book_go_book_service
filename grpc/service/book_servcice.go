package service

import (
	"book/book_go_book_service/config"
	"book/book_go_book_service/genproto/book_service"
	"book/book_go_book_service/grpc/client"
	"book/book_go_book_service/pkg/logger"
	"book/book_go_book_service/storage"
)

type bookService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	book_service.UnimplementedBookServiceServer
}

func NewBookService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *bookService {
	return &bookService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}
