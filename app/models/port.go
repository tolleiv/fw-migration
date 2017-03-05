package models

import "github.com/jinzhu/gorm"

type Port struct {
	gorm.Model
	Name string
	Num  string
}
