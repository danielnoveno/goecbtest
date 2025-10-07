package models

import (
	"gorm.io/gorm"
)

type Theme struct {
	gorm.Model
	Id				uint   `gorm:"primaryKey;autoIncrement" json:"theme_id"`
	Nama 			string `gorm:"size:20"`
	Keterangan 		string `gorm:"type:text"`
	Layout 			string `gorm:"size:20"`
	Folder 			string `gorm:"size:255"`
	SkinCssstring 	string `gorm:"size:255"`
	AppJs 			string `gorm:"size:255"`
	ScreenshotUrl	string `gorm:"size:255"`
	// CreateAt		time.Time
	// UpdatedAt		time.Time
}