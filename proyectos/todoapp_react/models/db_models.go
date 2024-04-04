package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"first_name"`
	LastName string `gorm:"not null" json:"last_name"`
	Email string `gorm:"not null;unique_index" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Tasks []Task `json:"tasks"`

}

type Task struct {
	gorm.Model

	Title string `gorm:"type:varchar(100);not null;unique_index" json:"title"`
	Description string `json:"description"`
	Done bool `gorm:"default:false" json:"done"`
	UserId uint `json:"user_id"`
	User User `gorm:"foreingKey:UserId"`
}