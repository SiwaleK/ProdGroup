// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

const TableNameSubdistrict = "subdistricts"

// Subdistrict mapped from table <subdistricts>
type Subdistrict struct {
	Subdistrictid     string `gorm:"column:subdistrictid;not null" json:"subdistrictid"`
	Zipcode           int32  `gorm:"column:zipcode;not null" json:"zipcode"`
	Subdistrictname   string `gorm:"column:subdistrictname;not null" json:"subdistrictname"`
	Subdistrictnameen string `gorm:"column:subdistrictnameen;not null" json:"subdistrictnameen"`
	Districtid        int32  `gorm:"column:districtid;not null" json:"districtid"`
}

// TableName Subdistrict's table name
func (*Subdistrict) TableName() string {
	return TableNameSubdistrict
}
