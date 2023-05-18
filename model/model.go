package models

import (
	"time"
)

type SKU_branch struct {
	Skuid      string    `json:"skuid"`
	Merchantid string    `json:"merchantid"`
	Branchid   string    `json:"branchid"`
	Price      float64   `json:"price"`
	Startdate  time.Time `json:"startdate"`
	Enddate    time.Time `json:"enddate"`
	Isactive   int8      `json:"isactive"`
}
type SKU struct {
	Skuid           string    `json:"skuid"`
	Barcodepos      string    `json:"barcodepos"`
	Productname     string    `json:"productname"`
	Brandid         int64     `json:"brandid"`
	Productgroupid  int64     `json:"productgroupid"`
	Productcatid    int64     `json:"productcatid"`
	Productsubcatid int64     `json:"productsubcatid"`
	Productsizeid   int64     `json:"productsizeid"`
	Productunit     int64     `json:"productunit"`
	Packsize        string    `json:"packsize"`
	Unit            int64     `json:"unit"`
	Banforpracharat int64     `json:"banforpracharat"`
	Isvat           int8      `json:"isvat"`
	Createby        string    `json:"createby"`
	Createdate      time.Time `json:"createdate"`
	Isactive        int8      `json:"isactive"`
	Merchantid      string    `json:"merchantid"`
	Mapsku          string    `json:"mapsku"`
	Isfixprice      int8      `json:"isfixprice"`
}
type ProductGroup struct {
	Prodgroupid int8   `json:"prodgroupid"`
	Th_name     string `json:"th_name"`
	En_name     string `json:"en_name"`
}
type ProductCategory struct {
	Prodcatid uint   `json:"prodcatid"`
	Th_name   string `json:"th_name"`
	En_name   string `json:"en_name"`
}

type ProductSubCategory struct {
	Prodsubcatid int8   `json:"prodgroupid"`
	Th_name      string `json:"th_name"`
	En_name      string `json:"en_name"`
}

type ProdCatProdSubCat struct {
	Prodcatid    int8 `json:"prodcatid"`
	Prodsubcatid int8 `json:"prodsubcatid"`
}

type ProductGroupWithSubCategories struct {
	Prodcatid      uint   `gorm:"column:prodcatid"`
	Th_name        string `gorm:"column:th_name"`
	En_name        string `gorm:"column:en_name"`
	Subcat_th_name string `gorm:"column:subcat_th_name"`
	Subcat_en_name string `gorm:"column:subcat_en_name"`
}
type ProductGroupWithCategories struct {
	// Prodcatid   uint   `gorm:"column:prodcatid"`
	Th_name     string `gorm:"column:th_name"`
	En_name     string `gorm:"column:en_name"`
	Cat_th_name string `gorm:"column:cat_th_name"`
	Cat_en_name string `gorm:"column:cat_en_name"`
}
type Brand struct {
	Brandid  uint   `json:"brandid" gorm:"primaryKey;autoIncrement"`
	Th_brand string `json:"th_brand" binding:"required"`
	En_brand string `json:"en_brand" binding:"required"`
}
