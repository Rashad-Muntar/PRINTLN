package models

type Job struct {
	Id int    `gorm:"primary key;autoIncrement" json:"id"`
	Description string
	File string `gorm:"not null"`
}