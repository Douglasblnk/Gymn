package models

import (
	"time"

	"gorm.io/gorm"
)

type Equipment struct {
	ID        int            `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	UID       string         `gorm:"<-:false;unique;default:gen_random_uuid();not null"`
	Name      string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"default:now();not null"`
	UpdatedAt time.Time      `gorm:"default:now();not null"`
	DeletedAt gorm.DeletedAt ``
}
