package model

import "gorm.io/gorm"

type UserTask struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	User User `gorm:"embedded"`
	Task Task `gorm:"embedded"`
}
