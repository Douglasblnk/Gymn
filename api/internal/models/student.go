package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        int            `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	UID       string         `gorm:"<-:false;unique;default:gen_random_uuid();not null"`
	FirstName string         `gorm:"not null"`
	LastName  string         `gorm:"not null"`
	Birth     *string        ``
	Color     *string        ``
	Code      string         `gorm:"not null;unique"`
	StartDate string         `gorm:"default:now();not null"`
	Weight    string         `gorm:"not null"`
	Height    string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"default:now();not null"`
	UpdatedAt time.Time      `gorm:"default:now();not null"`
	DeletedAt gorm.DeletedAt ``
}
