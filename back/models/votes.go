package models

import (
	"github.com/jinzhu/gorm"
)

// Vote a struct relative to a votes
type Vote struct {
	gorm.Model
	ID          uint   `gorm:"primary_key"json:"-"`
	UUID        string `gorm:"unique_index"json:"uuid"`
	Title       string `gorm:"size:255"json:"Title"`
	Description string `gorm:"size:1023"json:"Description"`
}
