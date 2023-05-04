package repository

import (
	"context"
	"database/sql"

	"github.com/SiwaleK/ProdGroup/model"
)

type PromotionRepository interface {
	GetPromotionByID(ctx context.Context, promotionID uint) (*model.Promotion, error)
}

type DBPromotionRepository struct {
	db *sql.DB
}

func NewDBPromotionRepository(db *sql.DB) *DBPromotionRepository {
	return &DBPromotionRepository{db: db}
}

func (r *DBPromotionRepository) GetPromotionByID(ctx context.Context, id uint) (*model.Promotion, error) {
	promotion := &model.Promotion{}
	err := r.db.QueryRowContext(ctx, "GetPromotion", id).Scan(&promotion.Promotionid, &promotion.Promotiontype, &promotion.Startdate, &promotion.Enddate, &promotion.Description, &promotion.Conditions)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return promotion, nil
}

// func (r *PromotionRepository) GetPromotions(ctx context.Context) ([]*model.Promotion, error) {
// 	rows, err := r.db.QueryContext(ctx, "SELECT * FROM promotions")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var promotions []*model.Promotion
// 	for rows.Next() {
// 		promotion := &model.Promotion{}
// 		err := rows.Scan(&promotion.Promotionid, &promotion.Promotiontype, &promotion.Startdate, &promotion.Enddate, &promotion.Description, &promotion.Conditions)
// 		if err != nil {
// 			return nil, err
// 		}
// 		promotions = append(promotions, promotion)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return promotions, nil
// }
