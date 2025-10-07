package models

import (
	"gorm.io/gorm"
)

type Ecbstation struct {
	gorm.Model
	Id			uint   `gorm:"primaryKey;autoIncrement" json:"ecbstation_id"`
	Ipaddress 	string `gorm:"size:20"`
	Location 	string `gorm:"size:50"`
	Mode 		string `gorm:"size:20"`
	Linetype	string `gorm:"size:20"`
	Lineids		string
	Lineactive	int64
	Ecbstate	string `gorm:"size:20"`
	Theme		string `gorm:"size:20"`
	Tacktime	int64
	Workcenter	string `gorm:"type:text"`
	Status		string `gorm:"size:20"`
	// CreateAt	time.Time
	// UpdatedAt	time.Time
}