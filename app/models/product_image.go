package models

import (
	"time"
)

type ProductImage struct {
	ID         string    `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Product    Product   `json:"product" gorm:"foreignKey:ProductID"`
	ProductID  string    `json:"product_id" gorm:"size:36;index"`
	Path       string    `json:"path" gorm:"type:text"`
	ExtraLarge string    `json:"extra_large" gorm:"type:text"`
	Large      string    `json:"large" gorm:"type:text"`
	Medium     string    `json:"medium" gorm:"type:text"`
	Small      string    `json:"small" gorm:"type:text"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime"`
}
