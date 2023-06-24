package models

import (
	// "github.com/gofrs/uuid"
)

type User struct {
	Id int    `gorm:"primary key;autoIncrement" json:"id"`
	Username string 
	Email string `gorm:"unique"`
	Password string `gorm:"not null"`
}


