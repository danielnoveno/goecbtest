package models

import (
	"gorm.io/gorm"
)

type Mastersfg struct {
	gorm.Model
	Id			uint   `gorm:"primaryKey;autoIncrement" json:"mastersfg_id"`
	Plant 		string `gorm:"size:10"`
	Mattype		string `gorm:"size:20"`
	Matdesc 	string `gorm:"size:50"`
	Sfgtype 	string `gorm:"size:20"`
	Sfgdesc 	string `gorm:"size:50"`
	// CreateAt	time.Time
	// UpdatedAt	time.Time
}