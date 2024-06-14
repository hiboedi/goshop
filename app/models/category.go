package models

import "time"

type Category struct {
	ID        string    `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	ParentID  string    `json:"parent_id" gorm:"size:36"`
	Section   Section   `json:"section" gorm:"foreignKey:SectionID"`
	SectionID string    `json:"section_id" gorm:"size:36;index"`
	Products  []Product `json:"products" gorm:"many2many:product_categories;"`
	Name      string    `json:"name" gorm:"size:100;"`
	Slug      string    `json:"slug" gorm:"size:100;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime"`
}
