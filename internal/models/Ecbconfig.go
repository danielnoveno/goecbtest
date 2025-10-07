package models

import (
	"gorm.io/gorm"
)

type EcbConfig struct {
	gorm.Model
	Id			uint   `gorm:"primaryKey;autoIncrement" json:"ecbconfig_id"`
	Section 	string `gorm:"size:20;default:''"`
	Variable 	string `gorm:"size:20;default:''"`
	Value 		string `gorm:"type:text;default:''"`
	Ordering	string `gorm:"size:20;default:'000'"`
	// CreateAt	time.Time
	// UpdatedAt	time.Time
}