package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);not null"`
	Username string `gorm:"type:varchar(50);unique;not null"`
	Password string `gorm:"type:text;not null"` // Store hashed passwords, not plain text
}
