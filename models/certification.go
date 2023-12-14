package models

import "gorm.io/gorm"

type Certification struct {
	gorm.Model
	Name           string `gorm:"type:varchar(100);not null"`
	Authority      string `gorm:"type:varchar(100);not null"`
	CertificateURL string `gorm:"type:varchar(100)"`
	ImageURL       string `gorm:"type:varchar(100)"`
}
