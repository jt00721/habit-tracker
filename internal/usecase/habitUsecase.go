package usecase

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/jt00721/habit-tracker/internal/domain"
	"github.com/jt00721/habit-tracker/internal/repository"
	"gorm.io/gorm"
)

type HabitUsecase struct {
	HabitRepo repository.HabitRepository
}

func (usecase *HabitUsecase) CreateHabit(habit *domain.Habit) error {
	if habit.Name == "" {
		return fmt.Errorf("habit name cannot be empty")
	}

	if !domain.IsValidFrequency(habit.Frequency) {
		return fmt.Errorf("invalid frequency type: %s", habit.Frequency)
	}

	if err := usecase.HabitRepo.Create(habit); err != nil {
		log.Println("Error creating habit:", err)
		return fmt.Errorf("failed to create habit")
	}

	return nil
}

func (usecase *HabitUsecase) GetAllHabits() ([]domain.Habit, error) {
	habits, err := usecase.HabitRepo.GetAll()
	if err != nil {
		log.Println("Error retrieving all habits:", err)
		return nil, fmt.Errorf("failed to get habits")
	}

	sort.Slice(habits, func(i, j int) bool {
		return habits[i].CurrentStreak > habits[j].CurrentStreak
	})
	return habits, nil
}

func (usecase *HabitUsecase) GetHabitByID(id uint) (*domain.Habit, error) {
	habit, err := usecase.HabitRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("habit not found")
		}
		log.Printf("Error retrieving habit with ID(%d): %v", id, err)
		return nil, fmt.Errorf("failed to retrieve habit")
	}
	return habit, nil
}

func (usecase *HabitUsecase) UpdateHabit(habit *domain.Habit) error {
	existingHabit, err := usecase.GetHabitByID(habit.ID)
	if err != nil {
		log.Println("Error retrieving habit while trying to update habit:", err)
		return fmt.Errorf("failed to retrieve existing habit")
	}

	if habit.Name == "" {
		return fmt.Errorf("habit name cannot be empty")
	}

	if !domain.IsValidFrequency(habit.Frequency) {
		return fmt.Errorf("invalid frequency type: %s", habit.Frequency)
	}

	existingHabit.Name = habit.Name
	existingHabit.Frequency = habit.Frequency

	err = usecase.HabitRepo.Update(existingHabit)
	if err != nil {
		log.Printf("Error updating habit with ID(%d): %v", habit.ID, err)
		return fmt.Errorf("failed to update habit")
	}

	log.Printf("Habit (%s) updated successfully", habit.Name)
	return nil
}

// func validateHabit(name, frequency string) bool {
// 	if name == "" {
// 		return fmt.Errorf("habit name cannot be empty")
// 	}

// 	validFrequencies := map[string]bool{"daily": true, "weekly": true, "monthly": true}
// 	return validFrequencies[frequency]
// }

func (usecase *HabitUsecase) DeleteHabit(id uint) error {
	habit, err := usecase.GetHabitByID(id)
	if err != nil {
		log.Println("Error: Tried to delete non-existing habit with ID:", id)
		return fmt.Errorf("habit not found")
	}

	err = usecase.HabitRepo.Delete(id)
	if err != nil {
		log.Println("Error deleting habit:", err)
		return fmt.Errorf("failed to delete habit")
	}

	log.Printf("Habit (%s) deleted successfully", habit.Name)
	return nil
}

func (usecase *HabitUsecase) MarkCompleted(id uint) error {
	habit, err := usecase.GetHabitByID(id)
	if err != nil {
		log.Println("Error fetching habit for completion", err)
		return fmt.Errorf("habit not found")
	}

	now := time.Now()

	if habit.LastCompletedAt == nil || (habit.Frequency == "daily" && habit.LastCompletedAt.Add(24*time.Hour).Before(now)) {
		habit.CurrentStreak = 1
	} else if habit.Frequency == "weekly" && habit.LastCompletedAt.Add(7*24*time.Hour).Before(now) {
		habit.CurrentStreak = 1
	} else {
		habit.CurrentStreak++
	}

	habit.LastCompletedAt = &now
	habit.TotalCompletions++

	if err := usecase.HabitRepo.SafeUpdate(habit); err != nil {
		log.Println("Error marking habit as completed:", err)
		return fmt.Errorf("failed to mark habit as complete")
	}

	log.Printf("Habit with ID(%d) marked as completed. Current streak: %d", id, habit.CurrentStreak)
	return nil
}

func (usecase *HabitUsecase) GetStreaks() ([]domain.Habit, error) {
	habits, err := usecase.HabitRepo.GetStreaks()
	if err != nil {
		log.Println("Error retrieving all habit streaks:", err)
		return nil, fmt.Errorf("failed to get all habit streaks")
	}
	return habits, nil
}
