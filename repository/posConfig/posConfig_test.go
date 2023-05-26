package posconfig_test

import (
	"errors"
	"testing"

	mocks "example.com/go-crud-api/db/mock"
	repo "example.com/go-crud-api/repository/posConfig"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetPosonfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()

		posClientID := "11111111-1111-1111-1111-111111111111"
		posClientUUID, _ := uuid.Parse(posClientID)
		expectedResults := []repo.PosConfig{{
			Posclientid:       "11111111-1111-1111-1111-111111111111",
			Branchid:          "456",
			Branchno:          "789",
			Branchname:        "Sample Branch",
			Branchaddress:     "123 Main Street",
			Accountname:       "Sample Account",
			Accountcode:       "AC001",
			Merchantid:        "M123",
			Merchantname:      "Sample Merchant",
			Taxid:             "T123",
			Rdnumber:          "RD001",
			Isdrawer:          true,
			Isbarcode:         false,
			Iscash:            true,
			Isqrcode:          false,
			Ispaotang:         true,
			Istongfah:         false,
			Iscoupon:          true,
			Sessiontype:       true,
			Barcodereadertype: false,
			Printertype:       false,
			Isactive:          true,
			Posrunning:        "POS001",
			Frposrunning:      "FRPOS001",
			Paymentmode:       1,
			IsVat:             1,
			Branchaddress2:    "456 Second Street",
			Branchsubdistrict: "Subdistrict",
			Branchdistrict:    "District",
			Branchprovince:    "Province",
			Branchzipcode:     "12345",
			Isinventory:       true,
			Isalertinventory:  false,
		},
		}
		rows := sqlmock.NewRows([]string{"Posclientid",
			"Branchid",
			"Branchno",
			"Branchname",
			"Branchaddress",
			"Accountname",
			"Accountcode",
			"Merchantid",
			"Merchantname",
			"Taxid",
			"Rdnumber",
			"Isdrawer",
			"Isbarcode",
			"Iscash",
			"Isqrcode",
			"Ispaotang",
			"Istongfah",
			"Iscoupon",
			"Sessiontype",
			"Barcodereadertype",
			"Printertype",
			"Isactive",
			"Posrunning",
			"Frposrunning",
			"Paymentmode",
			"IsVat",
			"Branchaddress2",
			"Branchsubdistrict",
			"Branchdistrict",
			"Branchprovince",
			"Branchzipcode",
			"Isinventory",
			"Isalertinventory"})
		for _, posconfig := range expectedResults {
			rows.AddRow(
				posconfig.Posclientid,
				posconfig.Branchid,
				posconfig.Branchno,
				posconfig.Branchname,
				posconfig.Branchaddress,
				posconfig.Accountname,
				posconfig.Accountcode,
				posconfig.Merchantid,
				posconfig.Merchantname,
				posconfig.Taxid,
				posconfig.Rdnumber,
				posconfig.Isdrawer,
				posconfig.Isbarcode,
				posconfig.Iscash,
				posconfig.Isqrcode,
				posconfig.Ispaotang,
				posconfig.Istongfah,
				posconfig.Iscoupon,
				posconfig.Sessiontype,
				posconfig.Barcodereadertype,
				posconfig.Printertype,
				posconfig.Isactive,
				posconfig.Posrunning,
				posconfig.Frposrunning,
				posconfig.Paymentmode,
				posconfig.IsVat,
				posconfig.Branchaddress2,
				posconfig.Branchsubdistrict,
				posconfig.Branchdistrict,
				posconfig.Branchprovince,
				posconfig.Branchzipcode,
				posconfig.Isinventory,
				posconfig.Isalertinventory,
			)
		}

		mock.ExpectQuery(".+").WillReturnRows(rows)
		posConfig := repo.NewPosConfigRepository()
		result, err := posConfig.GetPosConfig(posClientUUID)
		assert.Nil(t, err)
		assert.Equal(t, expectedResults, result)

	})

	t.Run("Fail", func(t *testing.T) {
		db, mock := mocks.Setup(t)
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
		posClientID := "11111111-1111-1111-1111-111111111111"
		posClientUUID, _ := uuid.Parse(posClientID)
		expectedError := errors.New("some database error ")
		mock.ExpectQuery(".+").WillReturnError(expectedError)
		posConfig := repo.NewPosConfigRepository()
		_, err := posConfig.GetPosConfig(posClientUUID)
		assert.Error(t, err)
		assert.Equal(t, expectedError, err)

	})
}
