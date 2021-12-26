package model

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  bool
}
