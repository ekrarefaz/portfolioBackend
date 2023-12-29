package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model           // Includes fields like ID, CreatedAt, UpdatedAt, DeletedAt
	Title       string   `gorm:"type:varchar(100);not null"`
	Description string   `gorm:"type:text;not null"`
	ProjectURL  string   `gorm:"type:varchar(100)"`
	GitHubURL   string   `gorm:"type:varchar(100)"`
	ImageURL    string   `gorm:"type:varchar(100)"`
	Badges      []string `gorm:"type:text" json:"-"`
}
