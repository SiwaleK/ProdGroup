package promotion

import (
	"time"

	"example.com/go-crud-api/db/database"
	"example.com/go-crud-api/db/db"
)

type PromotionRepository interface {
	CreatePromotion(arg []db.Promotion) error
	CreatePromotionAppliedItem(arg []db.PromotionAppliedItemsID) error
}

type promotionRepository struct {
}

func NewPromotionRepository() PromotionRepository {
	return &promotionRepository{}
}

type AppliedItems struct {
	PromotiondetailID string `json:"promotiondetail_ID"`
	Skuid             string `json:"skuid"`
}
type PostPromotionDiscountRequest struct {
	Promotionid    string            `json:"id"`
	Promotiontitle string            `json:"title"`
	Promotiontype  int32             `json:"promotiontype"`
	Startdate      time.Time         `json:"startDate"`
	Enddate        time.Time         `json:"endDate"`
	Description    string            `json:"text"`
	Condition      ConditionDiscount `json:"condition"`
	AppliedItemsID []AppliedItems    `json:"appliedItemsId"`
}

type ConditionDiscount struct {
	Discount int32
}

type PostPromotionAFREEBRequest struct {
	Promotionid    string          `json:"id"`
	Promotiontitle string          `json:"title"`
	Promotiontype  int32           `json:"promotiontype"`
	Startdate      time.Time       `json:"startdate"`
	Enddate        time.Time       `json:"enddate"`
	Description    string          `json:"text"`
	Condition      ConditionAFREEB `json:"condition"`
	AppliedItemsID []AppliedItems  `json:"appliedItemsId"`
}
type ConditionAFREEB struct {
	MinimumAmountToEnable int
	FreeAmount            int
	PremiumItemsId        []string
}

type ConditionStepPurchase struct {
	SpecialPriceAtXItemConditionDetail []SpecialPriceAtXItemConditionDetail `json:"specialPriceAtXItemConditionDetail"`
}

type SpecialPriceAtXItemConditionDetail struct {
	MinimumItemToEnable int `json:"minimumItemToEnable"`
	Discount            int `json:"discount"`
}

type PostPromotionStepPurchaseRequest struct {
	Promotionid    string                `json:"id"`
	Promotiontitle string                `json:"title"`
	Promotiontype  int32                 `json:"promotiontype"`
	Startdate      time.Time             `json:"startdate"`
	Enddate        time.Time             `json:"enddate"`
	Description    string                `json:"text"`
	Condition      ConditionStepPurchase `json:"condition"`
	AppliedItemsID []AppliedItems        `json:"appliedItemsId"`
}

func (r *promotionRepository) CreatePromotion(promotions []db.Promotion) error {
	if err := database.DB.Table("promotion").
		Create(&promotions).Error; err != nil {
		return err
	}
	return nil
}

func (r *promotionRepository) CreatePromotionAppliedItem(items []db.PromotionAppliedItemsID) error {
	if err := database.DB.Table("promotion_applied_items_id").
		Create(&items).Error; err != nil {
		return err
	}
	return nil
}
