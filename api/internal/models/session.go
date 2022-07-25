package models

type Session struct {
	ID           int    `gorm:"<-:false;primaryKey;autoIncrement;not null"`
	UID          string `gorm:"<-:false;unique;default:gen_random_uuid();not null"`
	UserID       int    `gorm:"not null"`
	RefreshToken string `gorm:"unique;not null"`
}
