package models

import (
	"time"

	"gorm.io/gorm"
)

type Ecbdata struct {
	gorm.Model
	Id			uint   `gorm:"primaryKey;autoIncrement" json:"ecbdata_id"`
	Tgl			time.Time `gorm:"default:NULL"`	
	Jam			time.Time `gorm:"default:NULL"`	
	Wc			string `gorm:"size:20;default:''"`
	Prdline		string `gorm:"size:20;default:''"`
	Ctgr		string `gorm:"default:'00:00:00'"`
	Sn 			string `gorm:"size:20;default:''"`
	Fgtype		string `gorm:"size:20;default:''"`
	Spc			string `gorm:"size:20;default:''"`
	Comptype	string `gorm:"size:20;default:''"`
	Compcode	string `gorm:"size:20;default:''"`
	Po			string `gorm:"size:20;default:''"`
	Sendsts		string `gorm:"size:20;default:''"`
	// CreateAt	time.Time
	// UpdateAt	time.Time
}