package models

import (
	"github.com/jinzhu/gorm"
)

// Vote a struct relative to a votes
type Vote struct {
	gorm.Model
	Uuid
	Title       string `gorm:"size:255"json:"title"`
	Description string `gorm:"size:1023"json:"description"`
	UserVotes   []User `gorm:"foreignkey:Uuid"`
}
