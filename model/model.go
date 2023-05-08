package model

import (
	"database/sql"
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
	Value int
	Title string
}

type Paymentconfig struct {
	IsCash              sql.NullInt64 `json:"is_cash"`
	IsQRCode            sql.NullInt64 `json:"is_qr_code"`
	PromtpayAccountName string        `json:"promtpay_account_name"`
	PromtpayNO          string        `json:"promtpay_no"`
	IsPaoTang           sql.NullInt64 `json:"is_pao_tang"`
	IsTongFah           sql.NullInt64 `json:"is_tong_fah"`
	IsCoupon            sql.NullInt64 `json:"is_coupon"`
	Printertype         int
	Printerconfig       []Printerconfig
}

type PaymentMethod struct {
	PaymentMethodID int    `gorm:"column:paymentmethodid;not null"`
	PaymentName     string `gorm:"column:paymentname;type:varchar(255)"`
}

func (Prodgroup) TableName() string {
	return "prodgroup" // specify the table name explicitly
}

type Posclient struct {
	PosClientID       sql.NullString `json:"pos_client_id"`
	BranchID          sql.NullString `json:"branch_id"`
	MerchantID        sql.NullString `json:"merchant_id"`
	RdNumber          sql.NullString `json:"rd_number"`
	IsDrawer          sql.NullInt64  `json:"is_drawer"`
	IsBarcode         sql.NullInt64  `json:"is_barcode"`
	IsCash            sql.NullInt64  `json:"is_cash"`
	IsQRCode          sql.NullInt64  `json:"is_qr_code"`
	IsPaoTang         sql.NullInt64  `json:"is_pao_tang"`
	IsTongFah         sql.NullInt64  `json:"is_tong_fah"`
	IsCoupon          sql.NullInt64  `json:"is_coupon"`
	SessionType       sql.NullInt64  `json:"session_type"`
	BarcodeReaderType sql.NullInt64  `json:"barcode_reader_type"`
	PrinterType       sql.NullInt64  `json:"printer_type"`
	IsActive          sql.NullInt64  `json:"is_active"`
	PosRunning        sql.NullString `json:"pos_running"`
	FrPosRunning      sql.NullString `json:"fr_pos_running"`
	PaymentMode       sql.NullInt64  `json:"payment_mode"`
}

type Branch struct {
	BranchID          sql.NullString `json:"branch_id"`
	MerchantID        sql.NullString `json:"merchant_id"`
	BranchNo          sql.NullString `json:"branch_no"`
	BranchName        sql.NullString `json:"branch_name"`
	BranchAddress     sql.NullString `json:"branch_address"`
	BranchEmail       sql.NullString `json:"branch_email"`
	AccountName       sql.NullString `json:"account_name"`
	AccountCode       sql.NullString `json:"account_code"`
	IsActive          int            `json:"is_active"`
	BranchAddress2    sql.NullString `json:"branch_address2"`
	BranchSubdistrict sql.NullString `json:"branch_subdistrict"`
	BranchDistrict    sql.NullString `json:"branch_district"`
	BranchProvince    sql.NullString `json:"branch_province"`
	BranchZipcode     sql.NullString `json:"branch_zipcode"`
	IsInventory       sql.NullInt64  `json:"is_inventory"`
	IsAlertInventory  sql.NullInt64  `json:"is_alert_inventory"`
}
