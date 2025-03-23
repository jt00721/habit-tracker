package usecase

import (
	"github.com/jt00721/habit-tracker/internal/domain"
)

// MockHabitRepo satisfies the HabitRepository interface
type MockHabitRepo struct {
	CreateFn     func(*domain.Habit) error
	GetAllFn     func() ([]domain.Habit, error)
	GetByIDFn    func(uint) (*domain.Habit, error)
	UpdateFn     func(*domain.Habit) error
	DeleteFn     func(uint) error
	SafeUpdateFn func(*domain.Habit) error
	GetStreaksFn func() ([]domain.Habit, error)
}

// Implement each method to call the corresponding function if set

func (m *MockHabitRepo) Create(h *domain.Habit) error {
	if m.CreateFn != nil {
		return m.CreateFn(h)
	}
	return nil
}

func (m *MockHabitRepo) GetAll() ([]domain.Habit, error) {
	if m.GetAllFn != nil {
		return m.GetAllFn()
	}
	return nil, nil
}

func (m *MockHabitRepo) GetByID(id uint) (*domain.Habit, error) {
	if m.GetByIDFn != nil {
		return m.GetByIDFn(id)
	}
	return nil, nil
}

func (m *MockHabitRepo) Update(h *domain.Habit) error {
	if m.UpdateFn != nil {
		return m.UpdateFn(h)
	}
	return nil
}

func (m *MockHabitRepo) Delete(id uint) error {
	if m.DeleteFn != nil {
		return m.DeleteFn(id)
	}
	return nil
}

func (m *MockHabitRepo) SafeUpdate(h *domain.Habit) error {
	if m.SafeUpdateFn != nil {
		return m.SafeUpdateFn(h)
	}
	return nil
}

func (m *MockHabitRepo) GetStreaks() ([]domain.Habit, error) {
	if m.GetStreaksFn != nil {
		return m.GetStreaksFn()
	}
	return nil, nil
}
