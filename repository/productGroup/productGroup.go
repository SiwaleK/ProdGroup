package productgroup

import (
	"example.com/go-crud-api/db/database"
)

type ProductGroupRepository interface {
	GetProductGroup() ([]Prodgroup, error)
	GetProductGroupByID(prodgroupID int) ([]Prodgroup, error)
}

type productGroupRepository struct{}

func NewProductGroupRepository() ProductGroupRepository {
	return &productGroupRepository{}
}

type Prodgroup struct {
	Prodgroupid int    `gorm:"column:prodgroupid;not null" json:"prodgroupid"`
	ThName      string `gorm:"column:th_name" json:"th_name"`
	EnName      string `gorm:"column:en_name" json:"en_name"`
}

func (r *productGroupRepository) GetProductGroup() ([]Prodgroup, error) {
	var productGroup []Prodgroup
	if err := database.DB.Table("prodgroup").
		Select("prodgroupid,th_name,en_name").
		Find(&productGroup).
		Error; err != nil {
		return nil, err
	}
	return productGroup, nil
}

func (r *productGroupRepository) GetProductGroupByID(prodgroupID int) ([]Prodgroup, error) {
	var productGroup []Prodgroup

	if err := database.DB.Table("prodgroup").
		Select("prodgroup.prodgroupid,prodgroup.th_name,prodgroup.en_name").
		Where("prodgroup.prodgroupid = ?", prodgroupID).
		Find(&productGroup).
		Error; err != nil {
		return nil, err
	}
	return productGroup, nil
}
