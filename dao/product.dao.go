package dao

import (
	"subscriptionManagement/model"

	"gorm.io/gorm"
)

type ProductDAO interface {
	Create(*model.Product) *gorm.DB
	Get(*[]model.Product) *gorm.DB
	GetById(uint, *model.Product) *gorm.DB
	Update(*model.Product) *gorm.DB
	Delete(uint) *gorm.DB
}
