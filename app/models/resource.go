package models

import "github.com/jinzhu/gorm"

type Resource struct {
	gorm.Model
	Name            string
	ResourceGroupID uint
	ResourceGroup   ResourceGroup
}
