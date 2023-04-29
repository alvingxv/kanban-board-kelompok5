package entity

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Fullname  string `gorm:"not null;column:full_name"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null;default:member"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
