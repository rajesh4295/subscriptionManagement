package model

import "gorm.io/gorm"

type Membership struct {
	gorm.Model
	TenantId string `json:"tenantId" binding:"required"`
	UserId   string `json:"userId" binding:"required"`
	IsActive bool   `json:"isActive" binding:"required"`
	Tenant   Tenant
	User     User
}
