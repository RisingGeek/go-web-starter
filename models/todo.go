package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Todo export
type Todo struct {
	gorm.Model
	ID      uuid.UUID `json:"ID" gorm:"primaryKey"`
	Message string    `json:"message"`
}
