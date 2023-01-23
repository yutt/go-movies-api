package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `json:"username" gorm:"unique,index"`
	Password  string
}
