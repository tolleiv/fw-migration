package models

import "github.com/jinzhu/gorm"

type Proto struct {
	gorm.Model
	Name string
}
