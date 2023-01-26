package service

import (
	"errors"
	"reflect"
	"testing"
	"unit-test/model"
	"unit-test/repository"

	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepository(ctrl)

	service := New(repo)

	t.Run("Create Success", func(t *testing.T) {

		repo.EXPECT().Create(model.Product{
			Code:  "12345",
			Price: 10000,
		}).Return(nil)

		err := service.Create(model.Product{
			Code:  "12345",
			Price: 10000,
		})

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Create Error Validation", func(t *testing.T) {

		err := service.Create(model.Product{
			Code:  "123456",
			Price: 10000,
		})

		if err == nil {
			t.Error(err)
		}
	})

	t.Run("Create Error Database", func(t *testing.T) {

		repo.EXPECT().Create(model.Product{
			Code:  "12345",
			Price: 10000,
		}).Return(errors.New("some error from database"))

		err := service.Create(model.Product{
			Code:  "12345",
			Price: 10000,
		})

		if err == nil {
			t.Error(err)
		}
	})

}

func TestFindByCode(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repository.NewMockRepository(ctrl)

	service := New(repo)

	t.Run("Find Success", func(t *testing.T) {

		repo.EXPECT().FindByCode("12345").Return(model.Product{
			Code:  "12345",
			Price: 10000,
		}, nil)

		product, err := service.FindByCode("12345")

		expectProduct := model.Product{
			Code:  "12345",
			Price: 10000,
		}

		if reflect.DeepEqual(product, expectProduct) == false {
			t.Error(product, expectProduct)
		}

		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Find Not Found", func(t *testing.T) {

		repo.EXPECT().FindByCode("12345").Return(model.Product{}, errors.New("record not found"))

		product, err := service.FindByCode("12345")

		expectProduct := model.Product{}

		if reflect.DeepEqual(product, expectProduct) == false {
			t.Error(product, expectProduct)
		}

		if err == nil {
			t.Error(err)
		}
	})

	t.Run("Find Error Validation", func(t *testing.T) {

		product, err := service.FindByCode("123456")

		expectProduct := model.Product{}

		if reflect.DeepEqual(product, expectProduct) == false {
			t.Error(product, expectProduct)
		}

		if err == nil {
			t.Error(err)
		}
	})
}
