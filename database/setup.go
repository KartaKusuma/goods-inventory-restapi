package database

import (
	"goods-inventory-restapi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/goods_inventory"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Goods{})

	DB = db
}
