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
	Description string         ``
	Repetitions string         `gorm:"not null"`
	Series      string         `gorm:"not null"`
	RestTime    string         ``
	Weight      float32        ``
	Cadence     string         ``
	Created_at  time.Time      `gorm:"default:now();not null"`
	Updated_at  time.Time      `gorm:"default:now();not null"`
	Deleted_at  gorm.DeletedAt ``
}
