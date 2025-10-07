package models

import (
	"time"

	"gorm.io/gorm"
)

type Ecbpo struct {
	gorm.Model
	Id			uint   `gorm:"primaryKey;autoIncrement" json:"ecbpo_id"`
	Worksenter 	string `gorm:"size:20;after:id"`
	Po 			string `gorm:"size:20"`
	Sn 			string `gorm:"size:20"`
	Ctype		string `gorm:"size:20"`
	Update_by	int64  `gorm:"unsigned"`
	EcbStatus 	string `gorm:"size:20"`
	CreateAt	time.Time
	UpdatedAt	time.Time
}