package models

import "time"

type Section struct {
	ID         string     `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name       string     `json:"name" gorm:"size:100;"`
	Slug       string     `json:"slug" gorm:"size:100;"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime"`
	Categories []Category `json:"categories"`
}
