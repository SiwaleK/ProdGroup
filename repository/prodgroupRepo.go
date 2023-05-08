package repository

import (
	"context"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
)

type ProdgroupRepository interface {
	GetProdgroup(ctx context.Context) ([]db.Prodgroup, error)
}

type DBProdgroupRepository struct {
	db *db.Queries
}

func NewProdgroupRepository(db *db.Queries) ProdgroupRepository {
	return &DBProdgroupRepository{
		db: db,
	}
}

func (r *DBProdgroupRepository) GetProdgroup(ctx context.Context) ([]db.Prodgroup, error) {
	dbProdgroups, err := r.db.GetProdgroup(ctx)
	if err != nil {
		return nil, err
	}
	prodgroups := make([]db.Prodgroup, len(dbProdgroups))
	for i, dbProdgroups := range dbProdgroups {
		prodgroups[i] = db.Prodgroup{
			Prodgroupid: dbProdgroups.Prodgroupid,
			ThName:      dbProdgroups.ThName,
			EnName:      dbProdgroups.EnName,
		}
	}

	return prodgroups, nil
}
