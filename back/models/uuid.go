package models

import (
    "github.com/jinzhu/gorm"
    uuid "github.com/satori/go.uuid"
    "time"
)

type Uuid struct {
    ID *uuid.UUID `gorm:"primary_key;unique;type:uuid;column:id"json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `sql:"index"json:"deleted_at"`
}

func (b *Base) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("ID", uuid.NewV4().String())
    return nil
}