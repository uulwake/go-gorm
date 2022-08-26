package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Outbound []Outbound `gorm:"constraint:OnDelete:CASCADE"`
	Name     string
	Qty      uint
	Weight   float32
}
