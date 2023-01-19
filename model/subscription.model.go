package model

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	PlanId   string `json:"planId" binding:"required"`
	TenantId string `json:"tenantId" binding:"required"`
	IsActive bool   `json:"isActive" binding:"required"`
	Plan     Plan
	Tenant   Tenant
}
