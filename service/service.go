package service

import (
	"errors"
	"unit-test/model"
	"unit-test/repository"
)

type Service interface {
	Create(product model.Product) error
	FindByCode(code string) (model.Product, error)
}

type ServiceImpl struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &ServiceImpl{
		repo: repo,
	}
}

func (s *ServiceImpl) Create(product model.Product) error {

	if len(product.Code) != 5 {
		return errors.New("product code must 5 character")
	}

	return s.repo.Create(product)
}

func (s *ServiceImpl) FindByCode(code string) (model.Product, error) {

	if len(code) != 5 {
		return model.Product{}, errors.New("product code must 5 character")
	}

	return s.repo.FindByCode(code)
}
