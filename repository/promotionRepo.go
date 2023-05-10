package repository

import (
	"context"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
)

// // type PromotionRepository struct {
// // 	queries *db.Queries
// // }

// // func NewDBPromotionRepository(queries *db.Queries) *PromotionRepository {
// // 	return &PromotionRepository{queries: queries}
// // }

// // func (r *PromotionRepository) GetPromotionByID(ctx context.Context, promotionID string) (*db.Promotion, error) {
// // 	dbPromotion, err := r.queries.GetPromotionByID(ctx, sql.NullString{
// // 		String: promotionID,
// // 		Valid:  promotionID != "",
// // 	})
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	promotions := &db.Promotion{
// // 		Promotionid: dbPromotion.Promotionid,
// // 		Startdate:   dbPromotion.Startdate,
// // 		Enddate:     dbPromotion.Enddate,
// // 		Description: dbPromotion.Description,
// // 		Condition:   dbPromotion.Condition,
// // 	}

// // 	return promotions, nil
// // }

type PromotionRepository interface {
	PostPromotion(ctx context.Context, arg db.PostPromotionParams) error
}

type DBPromotionRepository struct {
	db *db.Queries
}

func NewPromotionRepository(db *db.Queries) PromotionRepository {
	return &DBPromotionRepository{
		db: db,
	}
}

func (repo *DBPromotionRepository) PostPromotion(ctx context.Context, arg db.PostPromotionParams) error {
	err := repo.db.PostPromotion(ctx, arg)
	return err
}
