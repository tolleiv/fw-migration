package models

import "github.com/jinzhu/gorm"

type SecurityZone struct {
	gorm.Model
	Name string
	VPN  bool
}
