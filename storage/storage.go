package storage

import (
	"book/book_go_book_service/genproto/book_service"
	"context"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
}

type BookRepoI interface {
	Create(ctx context.Context, req *book_service.CreateBook) (resp *book_service.BookPrimaryKey, err error)
	Get(ctx context.Context, req *book_service.BookPrimaryKey) (resp *book_service.Book, err error)
	GetList(ctx context.Context, req *book_service.GetListBookRequest) (resp *book_service.GetListBookResponse, err error)
	Update(ctx context.Context, req *book_service.UpdateBook) (rowsAffected int64, err error)
	PatchUpdate(ctx context.Context, req *book_service.UpdatePatchBook) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *book_service.BookPrimaryKey) (rowsAffected int64, err error)
}
