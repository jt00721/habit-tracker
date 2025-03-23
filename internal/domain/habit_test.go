package domain

import (
	"testing"
	"time"
)

func TestHabitInitialisation(t *testing.T) {
	now := time.Now()
	habit := Habit{
		ID:               1,
		Name:             "Exercise",
		Frequency:        string(Daily),
		CurrentStreak:    5,
		LastCompletedAt:  &now,
		TotalCompletions: 10,
	}

	if habit.Name != "Exercise" {
		t.Errorf("Expected Name to be 'Exercise', got %s", habit.Name)
	}
}
