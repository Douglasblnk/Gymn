package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int            `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	Name        string         `gorm:"not null"`
	Email       string         `gorm:"unique;not null"`
	Password    string         `gorm:"not null"`
	Is_personal bool           `gorm:"default:FALSE;not null"`
	Photo       *string        ``
	CreatedAt   time.Time      `gorm:"default:now();not null"`
	UpdatedAt   time.Time      `gorm:"default:now();not null"`
	DeletedAt   gorm.DeletedAt ``
}
