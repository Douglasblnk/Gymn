package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            int              `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	UID           string           `gorm:"<-:false;unique;default:gen_random_uuid();not null"`
	Name          string           `gorm:"not null"`
	Email         string           `gorm:"unique;not null"`
	Password      string           `gorm:"not null"`
	Is_personal   *bool            `gorm:"default:FALSE;not null"`
	Photo         *string          ``
	Students      []*Student       `gorm:"foreignKey:UserID"`
	TrainingSheet []*TrainingSheet `gorm:"foreignKey:UserID"`
	Workout       []*Workout       `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time        `gorm:"default:now();not null"`
	UpdatedAt     time.Time        `gorm:"default:now();not null"`
	DeletedAt     gorm.DeletedAt   ``
}
