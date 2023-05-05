// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"database/sql"
	"encoding/json"
	"time"
)

type PaymentMethod struct {
	Paymentmethodid int32          `json:"paymentmethodid"`
	Paymentname     sql.NullString `json:"paymentname"`
}

type Prodgroup struct {
	Prodgroupid int32          `json:"prodgroupid"`
	ThName      sql.NullString `json:"th_name"`
	EnName      sql.NullString `json:"en_name"`
}

type Promotion struct {
	Promotionid   sql.NullString  `json:"promotionid"`
	Promotiontype int32           `json:"promotiontype"`
	Startdate     time.Time       `json:"startdate"`
	Enddate       time.Time       `json:"enddate"`
	Description   sql.NullString  `json:"description"`
	Conditions    json.RawMessage `json:"conditions"`
}

type PromotionAppliedItemsID struct {
	Promotiondetailid sql.NullString `json:"promotiondetailid"`
	Promotionid       sql.NullString `json:"promotionid"`
	Skuid             sql.NullString `json:"skuid"`
}