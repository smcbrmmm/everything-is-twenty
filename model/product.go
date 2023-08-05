package model

type Product struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	ProductName string `json:"productName" gorm:"primaryKey" gorm:"unique" binding:"required"`
	Quantity    int    `json:"quantity" binding:"gte=0"`
}
