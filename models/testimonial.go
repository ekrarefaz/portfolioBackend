package models

import "gorm.io/gorm"

type Testimonial struct {
	gorm.Model
	Author      string `gorm:"type:varchar(100);not null"`
	AuthorEmail string `gorm:"type:varchar(100);not null"`
	Content     string `gorm:"type:text;not null"`
	AuthorImage string `gorm:"type:varchar(100)"`
}
