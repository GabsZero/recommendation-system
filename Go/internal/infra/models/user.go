package models

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Movies    []Movie `gorm:"many2many:user_movies;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
