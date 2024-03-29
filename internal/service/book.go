package service

import (
	"github.com/Windmill787-golang/junior-test/internal/entities"
	"github.com/Windmill787-golang/junior-test/internal/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo}
}

func (s *BookService) GetBook(id int) (*entities.Book, error) {
	return s.repo.GetBook(id)
}

func (s *BookService) GetBooks() ([]*entities.Book, error) {
	return s.repo.GetBooks()
}

func (s *BookService) CreateBook(book entities.Book) (int, error) {
	return s.repo.CreateBook(book)
}

func (s *BookService) UpdateBook(book entities.Book) error {
	return s.repo.UpdateBook(book)
}

func (s *BookService) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}
