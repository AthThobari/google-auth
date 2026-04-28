package model

import "time"

type User struct {
	ID        string `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;not null"`
	Img       string
	Name      string
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time
}
