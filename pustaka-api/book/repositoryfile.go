package book

import (
	"errors"
	"fmt"
)

type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (r *fileRepository) FindAll() ([]Book, error) {
	var books []Book

	fmt.Println("Find All")

	return books, errors.New("")
}

func (r *fileRepository) FindByID(ID int) (Book, error) {
	var book Book

	fmt.Println("FindByID")

	return book, nil
}

func (r *fileRepository) Create(book Book) (Book, error) {
	fmt.Println("Create")

	return book, nil
}
