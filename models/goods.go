package models

type Goods struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
	Stock       int    `json:"stock" gorm:"type:integer"`
}
