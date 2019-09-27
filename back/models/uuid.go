package models

import (
    "github.com/jinzhu/gorm"
    uuid "github.com/satori/go.uuid"
	"fmt"
)

type Uuid struct {
    UUID *uuid.UUID `gorm:"unique;type:uuid;column:uuid"json:"uuid"`
}

func (b *Uuid) BeforeCreate(scope *gorm.Scope) error {
	u, err := uuid.NewV4()
	scope.SetColumn("UUID", u.String())
	if err != nil {
		fmt.Println(err)
	}
	return nil
}