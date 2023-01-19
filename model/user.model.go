package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"name" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	UserName  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	IsActive  bool   `json:"isActive" binding:"required"`
}
