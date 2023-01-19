package model

type Product struct {
	// gorm.Model
	CommonModelFields
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Url      string `json:"url" binding:"required"`
	IsActive *bool  `gorm:"not null" json:"isActive" binding:"required"`
}
