package models

import "github.com/jinzhu/gorm"

type Action struct {
	gorm.Model
	SourceID      uint
	Source        ResourceGroup
	DestinationID uint
	Destination   ResourceGroup
	ApplicationID uint
	Application   Application
	Action        string
}
