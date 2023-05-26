package productgroup_test

import (
	"errors"
	"testing"

	mocks "example.com/go-crud-api/db/mock"

	repo "example.com/go-crud-api/repository/productGroup"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetProductGroup(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
		expectedResults := []repo.Prodgroup{
			{
				Prodgroupid: 1,
				ThName:      "ไม่ได้จัดหมวดหมู่",
				EnName:      "No Group",
			},
			{
				Prodgroupid: 2,
				ThName:      "เครื่องดื่มแอลกอฮอล์",
				EnName:      "Alcoholic Beverage",
			},
		}
		rows := sqlmock.NewRows([]string{"prodgroupid", "th_name", "en_name"})
		for _, productGroup := range expectedResults {
			rows.AddRow(productGroup.Prodgroupid, productGroup.ThName, productGroup.EnName)
		}
		mock.ExpectQuery(".+").WillReturnRows(rows)
		productgroup := repo.NewProductGroupRepository()
		result, err := productgroup.GetProductGroup()
		assert.Nil(t, err)
		assert.Equal(t, expectedResults, result)
	})
	t.Run("failure", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
		expectedError := errors.New("some database err")
		mock.ExpectQuery(".+").WillReturnError(expectedError)
		productgroup := repo.NewProductGroupRepository()
		_, err := productgroup.GetProductGroup()
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestGetProductGroupByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		prodgroupid := 1

		expectedResults := []repo.Prodgroup{
			{
				Prodgroupid: 1,
				ThName:      "ไม่ได้จัดหมวดหมู่",
				EnName:      "No Group",
			},
		}
		rows := sqlmock.NewRows([]string{"prodgroupid", "th_name", "en_name"})
		for _, productGroup := range expectedResults {
			rows.AddRow(productGroup.Prodgroupid, productGroup.ThName, productGroup.EnName)
		}
		mock.ExpectQuery(".+").WillReturnRows(rows)
		productgroup := repo.NewProductGroupRepository()
		result, err := productgroup.GetProductGroupByID(prodgroupid)
		assert.Nil(t, err)
		assert.Equal(t, expectedResults, result)
	})

	t.Run("failure", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
		prodgroupid := 1
		expectedError := errors.New("some database err")
		mock.ExpectQuery(".+").WillReturnError(expectedError)
		productgroup := repo.NewProductGroupRepository()
		_, err := productgroup.GetProductGroupByID(prodgroupid)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
