package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);not null;unique"`
	Password  string `gorm:"type:varchar(100);not null"`
	IsActive  bool   `gorm:"type:boolean;default:true"`
}

func (u *User) TableName() string {
	return "users"
}
