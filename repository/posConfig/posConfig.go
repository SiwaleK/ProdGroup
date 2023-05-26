package posconfig

import (
	"example.com/go-crud-api/db/database"
	"github.com/google/uuid"
)

type PosConfigRepository interface {
	GetPosConfig(posClientID uuid.UUID) ([]PosConfig, error)
}

type posConfigRepository struct {
}

func NewPosConfigRepository() PosConfigRepository {
	return &posConfigRepository{}
}

type GetPosClientReq struct {
	PosClientID uuid.UUID `json:"posClientID"`
}

type PosConfig struct {
	Posclientid       string `gorm:"column:posclientid" json:"posClientID"`
	Branchid          string `gorm:"column:branchid" json:"branchID"`
	Branchno          string `gorm:"column:branchno" json:"branchNo"`
	Branchname        string `gorm:"column:branchname" json:"branchName"`
	Branchaddress     string `gorm:"column:branchaddress" json:"branchAddress"`
	Accountname       string `gorm:"column:accountname" json:"accountName"`
	Accountcode       string `gorm:"column:accountcode" json:"accountCode"`
	Merchantid        string `gorm:"column:merchantid" json:"merchantID"`
	Merchantname      string `gorm:"column:merchantname" json:"merchantName"`
	Taxid             string `gorm:"column:taxid" json:"taxID"`
	Rdnumber          string `gorm:"column:rdnumber" json:"rdNumber"`
	Isdrawer          bool   `gorm:"column:isdrawer" json:"isDrawer"`
	Isbarcode         bool   `gorm:"column:isbarcode" json:"isBarcode"`
	Iscash            bool   `gorm:"column:iscash" json:"isCash"`
	Isqrcode          bool   `gorm:"column:isqrcode" json:"isQrcode"`
	Ispaotang         bool   `gorm:"column:ispaotang" json:"isPaotang"`
	Istongfah         bool   `gorm:"column:istongfah" json:"isTongfah"`
	Iscoupon          bool   `gorm:"column:iscoupon" json:"isCoupon"`
	Sessiontype       bool   `gorm:"column:sessiontype" json:"sessionType"`
	Barcodereadertype bool   `gorm:"column:barcodereadertype" json:"barcodeReaderType"`
	Printertype       bool   `gorm:"column:printertype" json:"printerType"`
	Isactive          bool   `gorm:"column:isactive;not null" json:"isActive"`
	Posrunning        string `gorm:"column:posrunning" json:"posRunning"`
	Frposrunning      string `gorm:"column:frposrunning" json:"frPosRunning"`
	Paymentmode       int16  `gorm:"column:paymentmode" json:"paymentMode"`
	IsVat             int16  `gorm:"column:isvat" json:"isVat"`
	Branchaddress2    string `gorm:"column:branchaddress2" json:"branchAddress2"`
	Branchsubdistrict string `gorm:"column:branchsubdistrict" json:"branchSubdistrict"`
	Branchdistrict    string `gorm:"column:branchdistrict" json:"branchDistrict"`
	Branchprovince    string `gorm:"column:branchprovince" json:"branchProvince"`
	Branchzipcode     string `gorm:"column:branchzipcode" json:"branchZipcode"`
	Isinventory       bool   `gorm:"column:isinventory" json:"isInventory"`
	Isalertinventory  bool   `gorm:"column:isalertinventory" json:"isAlertinventory"`
}

func (r *posConfigRepository) GetPosConfig(posClientID uuid.UUID) ([]PosConfig, error) {
	var posConfig []PosConfig
	if err := database.DB.Table("posclient").
		Select("posclient.posclientid,branch.branchid,branch.branchNo,branch.branchName,branch.branchAddress,branch.accountName,branch.accountCode,merchant.merchantid,merchant.merchantname,merchant.taxid,posclient.iscash, posclient.iscoupon, posclient.ispaotang, posclient.istongfah, posclient.iscoupon, posclient.sessiontype, posclient.barcodereadertype , posclient.printertype , posclient.isactive, posclient.posrunning , posclient.frposrunning,posclient.rdnumber,branch.branchaddress2,branch.branchsubdistrict,branch.branchdistrict,branch.branchprovince,branch.branchzipcode,branch.isinventory,branch.isalertinventory,posclient.paymentmode").
		Joins("INNER JOIN branch on posclient.branchid = branch.branchid").
		Joins("INNER JOIN merchant ON branch.merchantid = merchant.merchantid").
		Where("posclient.posclientid=?", posClientID).
		Find(&posConfig).
		Error; err != nil {
		return []PosConfig{}, err
	}
	return posConfig, nil

}
