package repository

import (
	"log"

	"github.com/jt00721/habit-tracker/internal/domain"
	"gorm.io/gorm"
)

type HabitRepository struct {
	DB *gorm.DB
}

func (repo *HabitRepository) Create(habit *domain.Habit) error {
	if err := repo.DB.Create(habit).Error; err != nil {
		log.Println("Error creating habit via repo:", err)
		return err
	}
	return nil
}

func (repo *HabitRepository) GetByID(id uint) (*domain.Habit, error) {
	var habit domain.Habit
	if err := repo.DB.First(&habit, id).Error; err != nil {
		log.Println("Error getting habit by ID via repo:", err)
		return nil, err
	}
	return &habit, nil
}

func (repo *HabitRepository) GetAll() ([]domain.Habit, error) {
	var habits []domain.Habit
	if err := repo.DB.Find(&habits).Error; err != nil {
		log.Println("Error getting all habits via repo:", err)
		return nil, err
	}
	return habits, nil
}

func (repo *HabitRepository) Update(habit *domain.Habit) error {
	if err := repo.DB.Save(habit).Error; err != nil {
		log.Println("Error updating habit via repo:", err)
		return err
	}
	return nil
}

func (repo *HabitRepository) Delete(id uint) error {
	if err := repo.DB.Delete(&domain.Habit{}, id).Error; err != nil {
		log.Println("Error deleting habit via repo:", err)
		return err
	}
	return nil
}
