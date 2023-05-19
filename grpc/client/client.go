package client

import (
	"book/book_go_book_service/config"
	"book/book_go_book_service/genproto/author_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	//BookService() book_service.BookServiceClient
	AuthorService() author_service.AuthorServiceClient
}

type grpcClients struct {
	// bookService book_service.BookServiceClient
	authorService author_service.AuthorServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connAuthorService, err := grpc.Dial(
		cfg.ServiceHost+cfg.ServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		authorService: author_service.NewAuthorServiceClient(connAuthorService),
	}, nil
}

func (g *grpcClients) AuthorService() author_service.AuthorServiceClient {
	return g.authorService
}
