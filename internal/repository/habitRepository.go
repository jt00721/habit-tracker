package repository

import (
	"github.com/jt00721/habit-tracker/internal/domain"
	"gorm.io/gorm"
)

type HabitRepository struct {
	DB *gorm.DB
}

func (repo *HabitRepository) Create(habit *domain.Habit) error {
	return repo.DB.Create(habit).Error
}

func (repo *HabitRepository) GetByID(id uint) (*domain.Habit, error) {
	var habit domain.Habit
	err := repo.DB.First(&habit, id).Error
	return &habit, err
}

func (repo *HabitRepository) GetAll() ([]domain.Habit, error) {
	var habits []domain.Habit
	err := repo.DB.Find(&habits).Error
	return habits, err
}

func (repo *HabitRepository) Update(habit *domain.Habit) error {
	return repo.DB.Save(habit).Error
}

func (repo *HabitRepository) Delete(id uint) error {
	return repo.DB.Delete(&domain.Habit{}, id).Error
}

func (repo *HabitRepository) SafeUpdate(habit *domain.Habit) error {
	tx := repo.DB.Begin()
	if err := tx.Save(habit).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (repo *HabitRepository) GetStreaks() ([]domain.Habit, error) {
	var habits []domain.Habit
	err := repo.DB.Where("current_streak > ?", 0).Order("current_streak DESC").Find(&habits).Error
	return habits, err
}
