package repository

import (
	"context"
	"database/sql"

	db "github.com/SiwaleK/ProdGroup/db/sqlc"
)

type PromotionRepository struct {
	queries *db.Queries
}

func NewDBPromotionRepository(queries *db.Queries) *PromotionRepository {
	return &PromotionRepository{queries: queries}
}

func (r *PromotionRepository) GetPromotionByID(ctx context.Context, promotionID string) (*db.Promotion, error) {
	dbPromotion, err := r.queries.GetPromotionByID(ctx, sql.NullString{
		String: promotionID,
		Valid:  promotionID != "",
	})
	if err != nil {
		return nil, err
	}

	promotions := &db.Promotion{
		Promotionid:   dbPromotion.Promotionid,
		Promotiontype: dbPromotion.Promotiontype,
		Startdate:     dbPromotion.Startdate,
		Enddate:       dbPromotion.Enddate,
		Description:   dbPromotion.Description,
		Conditions:    dbPromotion.Conditions,
	}

	return promotions, nil
}
