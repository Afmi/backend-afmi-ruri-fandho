package model

import "time"

type Cart struct {
	Id           string     `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	OwnerId      string     `gorm:"column:owner_id" json:"owner_id"`
	ProductId    string     `gorm:"column:product_id" json:"product_id"`
	Price        float64    `gorm:"column:price" json:"price"`
	TotalPrice   float64    `gorm:"column:total_price" json:"total_price"`
	FreeShipping *bool      `gorm:"column:free_shipping" json:"free_shipping"`
	Discount     string     `gorm:"column:discount" json:"discount"`
	IsDeleted    *bool      `gorm:"column:is_deleted" json:"is_deleted"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	DeletedBy    string     `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedAt    *time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy    string     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy    string     `gorm:"column:updated_by" json:"updated_by"`
}
