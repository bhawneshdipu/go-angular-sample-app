package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  bool
}
