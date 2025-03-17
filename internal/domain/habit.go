package domain

import "time"

type Habit struct {
	ID               uint   `gorm:"primaryKey"`
	Name             string `gorm:"not null"`
	Frequency        string `gorm:"not null"`
	CurrentStreak    int
	LastCompletedAt  *time.Time // Use pointer to handle null values
	TotalCompletions int
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}
