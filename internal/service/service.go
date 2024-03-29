package service

import (
	"github.com/Windmill787-golang/junior-test/internal/entities"
	"github.com/Windmill787-golang/junior-test/internal/repository"
)

type Book interface {
	GetBook(id int) (*entities.Book, error)
	GetBooks() ([]*entities.Book, error)
	CreateBook(book entities.Book) (int, error)
	UpdateBook(book entities.Book) error
	DeleteBook(id int) error
}

type Auth interface {
	CreateUser(user entities.User) (int, error)
	GenerateToken(use entities.User) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Book
	Auth
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repo.Book),
		Auth: NewAuthService(repo.Auth),
	}
}
