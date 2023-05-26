package mocks

import (
	"errors"

	repo "example.com/go-crud-api/repository/posConfig"
	"github.com/google/uuid"
)

type MockPosConfigRepository struct {
	MockGetPosConfig func(posClientID uuid.UUID) ([]repo.PosConfig, error)
}

func (m *MockPosConfigRepository) GetPosConfig(posClientID uuid.UUID) ([]repo.PosConfig, error) {
	if m.MockGetPosConfig != nil {
		return m.MockGetPosConfig(posClientID)
	}
	return nil, errors.New("MockGetPosConfig is not implemented")
}

func NewMockPosConfigRepository() *MockPosConfigRepository {
	return &MockPosConfigRepository{
		MockGetPosConfig: func(posClientID uuid.UUID) ([]repo.PosConfig, error) {
			return []repo.PosConfig{
				{
					Posclientid:       "1111-1111-1111",
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
			}, nil
		},
	}
}
