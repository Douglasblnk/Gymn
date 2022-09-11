package models

import (
	"time"

	"gorm.io/gorm"
)

type Workout struct {
	ID          int            `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	UID         string         `gorm:"<-:false;unique;default:gen_random_uuid();not null"`
	UserID      int            `gorm:"not null"`
	Exercise    string         `gorm:"not null"`
	Description string         `gorm:"not null"`
	Repetitions string         `gorm:"not null"`
	Series      int            `gorm:"not null"`
	RestTime    int            `gorm:"not null"`
	Weight      float32        `gorm:"not null"`
	Cadence     string         `gorm:"not null"`
	Created_at  time.Time      `gorm:"default:now();not null"`
	Updated_at  time.Time      `gorm:"default:now();not null"`
	Deleted_at  gorm.DeletedAt ``
}
