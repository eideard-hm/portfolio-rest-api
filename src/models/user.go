package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null;size:50" json:"first_name"`
	LastName  string `gorm:"not null;size:50" json:"last_name"`
	Email     string `gorm:"not null;size:50;unique_index" json:"email"`
}
