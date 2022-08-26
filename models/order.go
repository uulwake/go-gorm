package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Outbounds        []Outbound `gorm:"constraint:OnDelete:CASCADE"`
	RecipientName    string
	RecipientAddress string
	Shipper          string
}
