package models

import "gorm.io/gorm"

type Outbound struct {
	gorm.Model
	ItemID  uint
	OrderID uint
	Qty     uint
}
