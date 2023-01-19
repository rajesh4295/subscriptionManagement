package service

import (
	"subscriptionManagement/model"

	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func (p *Product) Create(product *model.Product) *gorm.DB {
	return p.DB.Save(product)
}

func (p *Product) Get(products *[]model.Product) *gorm.DB {
	return p.DB.Find(products)
}

func (p *Product) GetById(id uint, product *model.Product) *gorm.DB {
	return p.DB.Where("id = ?", id).Limit(1).Find(product)
}

func (p *Product) Update(product *model.Product) *gorm.DB {
	return p.DB.Where("id = ?", product.ID).Updates(product)
}

func (p *Product) Delete(id uint) *gorm.DB {
	return p.DB.Where("id = ?", id).Limit(1).Delete(&model.Product{})
}
