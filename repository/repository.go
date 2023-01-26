package repository

import (
	"unit-test/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(product model.Product) error
	FindByCode(code string) (model.Product, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) Create(product model.Product) error {
	tx := r.db.Create(&product)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *RepositoryImpl) FindByCode(code string) (model.Product, error) {
	product := model.Product{}

	tx := r.db.Where("code", code).First(&product)
	if tx.Error != nil {
		return product, tx.Error
	}

	return product, nil
}
