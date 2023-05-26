package mock

import (
	"errors"

	repo "example.com/go-crud-api/repository/productGroup"
)

type MockProductGroupRepository struct {
	MockGetProductGroup     func() ([]repo.Prodgroup, error)
	MockGetProductGroupByID func(prodgroupID int) ([]repo.Prodgroup, error)
}

func (m *MockProductGroupRepository) GetProductGroup() ([]repo.Prodgroup, error) {
	if m.MockGetProductGroup != nil {
		return m.MockGetProductGroup()
	}
	return nil, errors.New("MockGetProductGroup is not implemented")

}

func (m *MockProductGroupRepository) GetProductGroupByID(prodgroupID int) ([]repo.Prodgroup, error) {
	if m.MockGetProductGroupByID != nil {
		return m.MockGetProductGroupByID(prodgroupID)
	}
	return nil, errors.New("MockGetProductGroupByID is not implemented")

}

func NewMockProductGroupRepository() *MockProductGroupRepository {
	return &MockProductGroupRepository{
		MockGetProductGroup: func() ([]repo.Prodgroup, error) {
			return []repo.Prodgroup{
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
			}, nil
		},
		MockGetProductGroupByID: func(prodgroupID int) ([]repo.Prodgroup, error) {
			return []repo.Prodgroup{
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
			}, nil
		},
	}
}
