package model

import (
	"encoding/json"
	"time"
)

type Prodgroup struct {
	Prodgroupid int    `json:"prodgroupid" gorm:"not null"`
	Th_name     string `json:"th_name" gorm:"null;type:varchar(255)"`
	En_name     string `json:"en_name" gorm:"null;type:varchar(255)"`
}

type Promotion struct {
	Promotionid   string          `json:"promotionid" gorm:"null;type:varchar(36)"`
	Promotiontype int             `json:"promotiontype" gorm:"not null"`
	Startdate     time.Time       `json:"startdate" gorm:"not null"`
	Enddate       time.Time       `json:"enddate" gorm:"not null"`
	Description   string          `json:"description" gorm:"null;type:varchar(1024)"`
	Conditions    json.RawMessage `json:"conditio" gorm:"not null"`
}

type Promotion_applied_items_id struct {
	Promotiondetailid string `json:"promotiondetailid" gorm:"not null;type:varchar(36)"`
	Promotionid       string `json:"promotionid" gorm:"not null;type:varchar(36)"`
	Skuid             string `json:"skuid" gorm:"not null;type:varchar(36)"`
}

//from swagger ui

type Printerconfig struct {
	value int
	title string
}

type Paymentconfig struct {
	iscash              bool
	isqrcode            bool
	promtpayaccountname string
	promtpayno          string
	ispaotang           bool
	istongfah           bool
	iscoupon            bool
	printertype         int
	printerconfig       []Printerconfig
}

type Paymentmethod struct {
	PaymentMethodID int    `gorm:"column:paymentmethodid;not null"`
	PaymentName     string `gorm:"column:paymentname;type:varchar(255)"`
}

func (Prodgroup) TableName() string {
	return "prodgroup" // specify the table name explicitly
}
