package models

type Job struct {
	Id          int `gorm:"primary key;autoIncrement" json:"id"`
	UserID      int
	Description string
	OnGoing     bool   `gorm:"not null"`
	Completed   bool   `gorm:"not null"`
	File        string `gorm:"not null"`
}
