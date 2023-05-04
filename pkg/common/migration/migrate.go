package main

import (
	"github.com/SiwaleK/ProdGroup/db/config"
	"github.com/SiwaleK/ProdGroup/model"
)

func main() {
	config.DB.AutoMigrate(&model.Prodgroup{})
}
