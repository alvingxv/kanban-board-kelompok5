package entity

import "time"

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string `gorm:"not null"`
	Task      []Task
	CreatedAt time.Time
	UpdatedAt time.Time
}
