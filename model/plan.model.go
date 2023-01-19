package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	ProductID uint           `json:"productId" binding:"required"`
	Name      string         `json:"name" binding:"required"`
	Price     uint           `json:"price" binding:"required"`
	Currency  string         `json:"currency" binding:"required"`
	Features  datatypes.JSON `json:"features" binding:"required"`
	IsActive  bool           `json:"isActive" binding:"required"`
	Product   Product
}
