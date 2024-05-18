package models

import (
	"time"
)

type Recommendation struct {
	ID        uint
	UserId    uint
	User      User
	MovieId   uint
	Movie     Movie
	CreatedAt time.Time
	UpdatedAt time.Time
}
