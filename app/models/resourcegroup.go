package models

import "github.com/jinzhu/gorm"

type ResourceGroup struct {
	gorm.Model
	Name           string
	SecurityZoneID uint
	SecurityZone   SecurityZone
}
