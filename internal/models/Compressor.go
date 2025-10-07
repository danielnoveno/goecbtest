package models

import (
	"gorm.io/gorm"
)

type Compressor struct{
	gorm.Model
	Id   		uint   `gorm:"primaryKey;autoIncrement" json:"compressor_id"`
	Ctype  		string `gorm:"size:4;default:''"`
	Merk 		string `gorm:"size:20;default:''"`
	CompType 	string `gorm:"size:20;default:''"`
	ItemCode 	string `gorm:"size:20;default:''"`
	ForceScan 	int64  `gorm:"default:1"`
	FamilyCode 	string `gorm:"size:20;default:''"`
	Status 		string `gorm:"size:20;default:''"`
	// CreateAt	time.Time
	// UpdatedAt	time.Time
}