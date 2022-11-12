package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

// example struct
type CategoryServices struct {
	Repository repository.CategoryRepository
}

// example 1
func (service CategoryServices) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
