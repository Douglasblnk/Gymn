package models

import (
	"time"

	"gorm.io/gorm"
)

type TrainingSheet struct {
	ID         int            `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	UID        string         `gorm:"<-:false;unique;default:gen_random_uuid();not null"`
	Name       string         `gorm:"not null"`
	Active     bool           `gorm:"not null;default:true"`
	Created_at time.Time      `gorm:"default:now();not null"`
	Updated_at time.Time      `gorm:"default:now();not null"`
	Deleted_at gorm.DeletedAt ``
}
