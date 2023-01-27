package model

import (
	"time"
)

type Film struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"unique"`
	Director    string    `json:"director"`
	ReleaseDate time.Time `json:"release_date"`
	Synopsis    string    `json:"synopsis" gorm:"varchar(1000)"`
	Genres      []Genre   `json:"genre " gorm:"many2many:film_genres;"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeleteAt    time.Time `json:"delete_at" gorm:"index"`
	CreatedBy   User      `json:"created_by" gorm:"foreignkey:user_id;association_foreignkey:ID;"`
	UserID      uint      `json:"uid" gorm:"index;not null;column:user_id"`
}

type Genre struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"delete_at" gorm:"index"`
}
