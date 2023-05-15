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
	//PostPromotion(ctx context.Context, arg db.PostPromotionParams, args []db.PostPromotionParams) error
	//PostPromotion2(ctx context.Context, arg db.PostPromotionParams) error
	PostPromotionAppliedItem(ctx context.Context, arg []db.PostPromotionAppliedParams) error
	//PostPromotion4(ctx context.Context, arg db.PostPromotionAppliedParams) error
	PostPromotion(ctx context.Context, arg db.PostPromotionTableParams) error
}

type DBPromotionRepository struct {
	db *db.Queries
}

func NewPromotionRepository(db *db.Queries) PromotionRepository {
	return &DBPromotionRepository{
		db: db,
	}
}

// func (repo *DBPromotionRepository) PostPromotion(ctx context.Context, arg db.PostPromotionParams, args []db.PostPromotionParams) error {
// 	// Use arg and args as needed
// 	// Call the underlying database implementation
// 	err := repo.db.PostPromotion(ctx, arg)
// 	return err
// }

// func (repo *DBPromotionRepository) PostPromotion2(ctx context.Context, arg db.PostPromotionParams) error {
// 	// Use arg as needed
// 	// Call the underlying database implementation
// 	err := repo.db.PostPromotion(ctx, arg)
// 	return err
// }

func (repo *DBPromotionRepository) PostPromotionAppliedItem(ctx context.Context, arg []db.PostPromotionAppliedParams) error {
	for _, a := range arg {
		err := repo.db.PostPromotionApplied(ctx, a)
		if err != nil {
			return err
		}
	}
	return nil
}

//	func (repo *DBPromotionRepository) PostPromotion4(ctx context.Context, arg db.PostPromotionAppliedParams) error {
//		// Use arg as needed
//		// Call the underlying database implementation
//		err := repo.db.PostPromotionApplied(ctx, arg)
//		return err
//	}
func (repo *DBPromotionRepository) PostPromotion(ctx context.Context, arg db.PostPromotionTableParams) error {
	// Use arg and args as needed
	// Call the underlying database implementation
	err := repo.db.PostPromotionTable(ctx, arg)
	return err
}
