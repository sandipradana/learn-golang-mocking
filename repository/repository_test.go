package repository

import (
	"errors"
	"reflect"
	"regexp"
	"testing"
	"unit-test/model"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {

	dbTest, mock, err := sqlmock.New()
	if err != nil {
		t.Error("Error Create SqlMock")
	}

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8"))

	gormTest, err := gorm.Open(mysql.New(mysql.Config{Conn: dbTest}), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	repo := New(gormTest)

	// Create Success
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `products` (`code`,`price`) VALUES (?,?)")).
		WithArgs("KODE-1", uint(1000)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Create(model.Product{
		Code:  "KODE-1",
		Price: 1000,
	})
	if err != nil {
		t.Error(err)
	}

	// Create Error
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `products` (`code`,`price`) VALUES (?,?)")).
		WithArgs("KODE-2", uint(2000)).WillReturnError(errors.New("id sudah ada"))
	mock.ExpectRollback()

	err = repo.Create(model.Product{
		Code:  "KODE-2",
		Price: 2000,
	})
	if err.Error() != "id sudah ada" {
		t.Error(err)
	}
}

func TestFindByCode(t *testing.T) {

	dbTest, mock, err := sqlmock.New()
	if err != nil {
		t.Error("Error Create SqlMock")
	}

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8"))

	gormTest, err := gorm.Open(mysql.New(mysql.Config{Conn: dbTest}), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	repo := New(gormTest)

	t.Run("Create Success", func(t *testing.T) {
		// Create Success
		expectRows := sqlmock.NewRows([]string{"code", "price"}).
			AddRow("KODE-1", uint(1000))
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `code` = ? ORDER BY `products`.`code` LIMIT 1")).
			WillReturnRows(expectRows)

		productExpect := model.Product{
			Code:  "KODE-1",
			Price: 1000,
		}
		product, err := repo.FindByCode("KODE-1")
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(product, productExpect) {
			t.Error("expect product tidak sama")
		}
	})

	t.Run("Create Error", func(t *testing.T) {
		// Create Error
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `products` WHERE `code` = ?  ORDER BY `products`.`code` LIMIT 1")).
			WillReturnError(errors.New("not found"))
		_, err = repo.FindByCode("KODE-2")
		if err.Error() != "not found" {
			t.Error(err)
		}
	})
}
