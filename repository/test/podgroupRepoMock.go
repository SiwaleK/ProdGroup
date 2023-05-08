package repository

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
	"github.com/stretchr/testify/mock"
)

type ProdgroupRepositoryMock struct {
	mock.Mock
}

func (r *ProdgroupRepositoryMock) GetProdgroup(ctx context.Context) (*db.Prodgroup, error) {
	args := r.Called(ctx)

	if prodgroup, ok := args.Get(0).(*db.Prodgroup); ok {
		return prodgroup, args.Error(1)
	}

	return nil, args.Error(1)
}

func ProdgroupMock() {
	repo := &ProdgroupRepositoryMock{}

	// Set up the mock response
	mockProdgroup := &db.Prodgroup{
		Prodgroupid: 1,
		ThName:      sql.NullString{String: "Product Group 1", Valid: true},
		EnName:      sql.NullString{String: "Group 1", Valid: true},
	}
	repo.On("GetProdgroup", mock.Anything).Return(mockProdgroup, nil)

	// Call the GetProdgroup method
	ctx := context.Background()
	prodgroup, err := repo.GetProdgroup(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Prodgroup:", prodgroup.Prodgroupid, prodgroup.ThName.String, prodgroup.EnName.String)
}
