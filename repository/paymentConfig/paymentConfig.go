package repo

import (
	"example.com/go-crud-api/db/database"
)

type paymentConfigRepository struct {
}

func NewPaymentConfigRepository() PaymentConfigRepository {
	return &paymentConfigRepository{}
}

type PromptPayData struct {
	AccountName string `json:"promptPayAccountName"`
	AccountCode string `json:"promptPayNo"`
}

type PaymentConfig struct {
	PromptPayData PromptPayData `json:"promptPayData"`
	Items         []Item        `json:"items"`
}

type Item struct {
	PaymentMethodID int32  `json:"paymentMethodID"`
	PaymentName     string `json:"paymentName"`
	IsEnable        *bool  `json:"isEnable"`
}

type GetPosClientMethodReq struct {
	PosClientID string `json:"posClientID"`
}

type PaymentMethod struct {
	Paymentmethodid int32  `gorm:"column:paymentmethodid;not null" json:"paymentmethodid"`
	Paymentname     string `gorm:"column:paymentname" json:"paymentname"`
}

type PosclientConfig struct {
	Iscash      bool   `gorm:"column:iscash" json:"iscash"`
	Isqrcode    bool   `gorm:"column:isqrcode" json:"isqrcode"`
	Ispaotang   bool   `gorm:"column:ispaotang" json:"ispaotang"`
	Istongfah   bool   `gorm:"column:istongfah" json:"istongfah"`
	Iscoupon    bool   `gorm:"column:iscoupon" json:"iscoupon"`
	Accountname string `gorm:"column:accountname" json:"accountname"`
	Accountcode string `gorm:"column:accountcode" json:"accountcode"`
}

type PaymentConfigRepository interface {
	GetPosClientConfig(posClientID string) ([]PosclientConfig, error)
	GetPaymentMethod() ([]PaymentMethod, error)
}

func (r *paymentConfigRepository) GetPosClientConfig(posClientID string) ([]PosclientConfig, error) {
	var posclientConfigResponse []PosclientConfig

	if err := database.DB.Table("posclient").
		Select("posclient.iscash, posclient.isqrcode, posclient.ispaotang, posclient.istongfah, posclient.iscoupon, branch.accountname, branch.accountcode").
		Joins("INNER JOIN branch ON branch.branchid = posclient.branchid").
		Where("posclient.posclientid = ?", posClientID).
		Find(&posclientConfigResponse).
		Error; err != nil {
		return nil, err
	}

	return posclientConfigResponse, nil
}

func (r *paymentConfigRepository) GetPaymentMethod() ([]PaymentMethod, error) {
	var PaymentMethodResult []PaymentMethod

	if err := database.DB.Table("payment_method").
		Select("payment_method.paymentmethodid, payment_method.paymentname").
		Find(&PaymentMethodResult).
		Error; err != nil {
		return []PaymentMethod{}, err

	}

	return PaymentMethodResult, nil
}
