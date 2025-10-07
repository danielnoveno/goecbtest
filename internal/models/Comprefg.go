package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Comprefg struct {
	gorm.Model
	Id			uint   `gorm:"primaryKey;autoIncrement" json:"comprefg_id"`
	Ctype 		string `gorm:"size:4;default:''"`
	Barcode 	string `gorm:"size:20;default:''"`
	Status 		string `gorm:"size:20;default:''"`
	// CreateAt	time.Time
	// UpdatedAt	time.Time
}

func (c *Comprefg) Validate() error{
	if c.Ctype == "" {
		return fmt.Errorf("ctype is required")
	}

	if c.Barcode == ""{
		return fmt.Errorf("barcode is required")
	}

	if c.Status == ""{
		return fmt.Errorf("status is required")
	}
	return nil
}

func NewComprefg(ctype, barcode, status string) * Comprefg {
	return &Comprefg{
		Ctype: ctype,
		Barcode: barcode,
		Status: status,
	}
}

func (c Comprefg) String() string {
	return fmt.Sprintf("Comprefg<%s %s %s>", c.Ctype, c.Barcode, c.Status)
}