package models

import (
	"gorm.io/gorm"
)

type Navigation struct {
	gorm.Model
	Id			uint	`gorm:"primaryKey;autoIncrement" json:"navigation_id"`
	Parent_id 	int64	`gorm:"default:NULL"`
	Icon 		string	`gorm:"size:100"`
	Title 		string	`gorm:"size:255"`
	Url 		string	`gorm:"size:255"`
	Mode 		int64	`gorm:"default:1"`
	Urutan 		int64	`gorm:"default:0"`
	// CreateAt	time.Time
	// UpdatedAt	time.Time
}