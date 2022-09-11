package models

import (
	"time"

	"gorm.io/gorm"
)

type TrainingSheet struct {
	ID         int            `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	UID        string         `gorm:"<-:false;unique;default:gen_random_uuid();not null"`
	UserID     int            `gorm:"not null"`
	Name       string         `gorm:"not null"`
	Active     bool           `gorm:"not null;default:true"`
	Student    []*Student     `gorm:"many2many:public.student_training_sheets"`
	Workout    []*Workout     `gorm:"many2many:public.training_sheets_workouts"`
	Created_at time.Time      `gorm:"default:now();not null"`
	Updated_at time.Time      `gorm:"default:now();not null"`
	Deleted_at gorm.DeletedAt ``
}
