package models

import (
	"gorm.io/gorm"
)

type Masterfg struct {
	gorm.Model
	Id				uint   `gorm:"primaryKey;autoIncrement" json:"masterfg_id"`
	Mattype 		string `gorm:"size:10"`
	Matdesc 		string `gorm:"size:50"`
	Fgtype 			string `gorm:"size:20"`
	Aging_tipes_id	int64 `gorm:"default:0"`
	Kdbar 			string `gorm:"size:20"`
	Warna 			string `gorm:"size:20"`
	Lotinv 			string `gorm:"size:20;"`
	Attrib 			string `gorm:"size:100"`
	Category 		string `gorm:"size:20;default:''"`
	// CreateAt		time.Time
	// UpdatedAt		time.Time
}