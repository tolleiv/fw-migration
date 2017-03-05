package models

import "github.com/jinzhu/gorm"

type Application struct {
	gorm.Model
	Name          string
	SourceID      uint
	Source        Port
	DestinationID uint
	Destination   Port
	ProtoID       uint
	Proto         Proto
}
