package repository

import (
	"belajar-golang-unit-test/entity"
)

// example interface
type CategoryRepository interface {
	FindById(id string) *entity.Category
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}
	category := arguments.Get(0).(entity.Category)
	return &category
}
