package repository

import (
	"context"
	"database/sql"

	"github.com/SiwaleK/ProdGroup/model"
)

type PromotionRepository interface {
	GetPromotionByID(ctx context.Context, promotionid string) (*model.Promotion, error)
}

type DBPromotionRepository struct {
	db *sql.DB
}

func NewDBPromotionRepository(db *sql.DB) *DBPromotionRepository {
	return &DBPromotionRepository{db: db}
}

func (r *DBPromotionRepository) GetPromotionByID(ctx context.Context, promotionid string) (*model.Promotion, error) {
	promotion := &model.Promotion{}
	err := r.db.QueryRowContext(ctx, "SELECT * FROM promotion WHERE promotionid = $1", promotionid).Scan(&promotion.Promotionid, &promotion.Promotiontype, &promotion.Startdate, &promotion.Enddate, &promotion.Description, &promotion.Conditions)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return promotion, nil
}

// type Queries struct {
// 	db *sql.DB
// }

// func New(db *sql.DB) *Queries {
// 	return &Queries{
// 		db: db,
// 	}
// }

// type PromotionRepositorySQL interface {
// 	GetPromotionByID(ctx context.Context, promotionid string) (*db.Promotion, error)
// }

// func NewPromotionRepositorySQL(db *sql.DB) PromotionRepositorySQL {
// 	return &DBPromotionRepositorySQL{
// 		db:      db,
// 		queries: New(db), // Instantiate the Queries struct
// 	}
// }

// type DBPromotionRepositorySQL struct {
// 	db      *sql.DB
// 	queries *db.Queries
// }

// func (r *DBPromotionRepositorySQL) GetPromotionByID(ctx context.Context, promotionid string) (*db.Promotion, error) {
// 	promotion, err := r.queries.GetPromotionByID(ctx, promotionid)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &promotion, nil
// }
