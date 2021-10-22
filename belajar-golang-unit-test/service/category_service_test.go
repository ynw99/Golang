package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// Program Mock
	categoryRepository.Mock.On("FindByID", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSucceed(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "laptop",
	}

	categoryRepository.Mock.On("FindByID", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
