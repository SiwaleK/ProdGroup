package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SiwaleK/ProdGroup/model"
)

type ProdgroupRepository interface {
	GetProdgroup(ctx context.Context) (*[]model.Prodgroup, error)
}

type DBProdgroupRepository struct {
	db *sql.DB
}

func NewProdgroupRepository(db *sql.DB) ProdgroupRepository {
	return &DBProdgroupRepository{
		db: db,
	}
}

func (r *DBProdgroupRepository) GetProdgroup(ctx context.Context) (*[]model.Prodgroup, error) {
	prodgroups := &[]model.Prodgroup{}
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM prodgroup")

	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pg model.Prodgroup
		err := rows.Scan(&pg.Prodgroupid, &pg.Th_name, &pg.En_name)
		if err != nil {
			fmt.Println("Error scanning rows:", err)
			return nil, err
		}
		*prodgroups = append(*prodgroups, pg)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil, err
	}

	return prodgroups, nil
}
