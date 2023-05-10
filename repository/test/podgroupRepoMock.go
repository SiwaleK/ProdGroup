package repository

import (
	"context"

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
