package models

import "gorm.io/gorm"

type About struct {
	gorm.Model
	Content   string `gorm:"type:text;not null"`
	ResumeURL string `gorm:"type:varchar(100)"`
}
