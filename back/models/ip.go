package models

import (
    "github.com/jinzhu/gorm"
)

type Ip struct {
	gorm.Model
	Adress string
	Blocked bool
	Count int
}