package model

import (
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	IsActive bool   `json:"isActive" binding:"required"`
}
