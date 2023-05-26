// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

const TableNameBranch = "branch"

// Branch mapped from table <branch>
type Branch struct {
	Branchid          string `gorm:"column:branchid" json:"branchid"`
	Merchantid        string `gorm:"column:merchantid" json:"merchantid"`
	Branchno          string `gorm:"column:branchno" json:"branchno"`
	Branchname        string `gorm:"column:branchname" json:"branchname"`
	Branchaddress     string `gorm:"column:branchaddress" json:"branchaddress"`
	Branchemail       string `gorm:"column:branchemail" json:"branchemail"`
	Accountname       string `gorm:"column:accountname" json:"accountname"`
	Accountcode       string `gorm:"column:accountcode" json:"accountcode"`
	Isactive          bool   `gorm:"column:isactive;not null" json:"isactive"`
	Branchaddress2    string `gorm:"column:branchaddress2" json:"branchaddress2"`
	Branchsubdistrict string `gorm:"column:branchsubdistrict" json:"branchsubdistrict"`
	Branchdistrict    string `gorm:"column:branchdistrict" json:"branchdistrict"`
	Branchprovince    string `gorm:"column:branchprovince" json:"branchprovince"`
	Branchzipcode     string `gorm:"column:branchzipcode" json:"branchzipcode"`
	Isinventory       bool   `gorm:"column:isinventory" json:"isinventory"`
	Isalertinventory  bool   `gorm:"column:isalertinventory" json:"isalertinventory"`
}

// TableName Branch's table name
func (*Branch) TableName() string {
	return TableNameBranch
}