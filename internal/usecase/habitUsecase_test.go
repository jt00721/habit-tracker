package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/jt00721/habit-tracker/internal/domain"
	"github.com/jt00721/habit-tracker/internal/usecase"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateHabit(t *testing.T) {
	tests := []struct {
		name        string
		input       domain.Habit
		mockCreate  func(*domain.Habit) error
		wantErr     bool
		errContains string
	}{
		{
			name: "valid habit",
			input: domain.Habit{
				Name:      "Read",
				Frequency: "daily",
			},
			mockCreate: func(h *domain.Habit) error {
				return nil
			},
			wantErr: false,
		},
		{
			name: "missing name",
			input: domain.Habit{
				Name:      "",
				Frequency: "daily",
			},
			wantErr:     true,
			errContains: "habit name cannot be empty",
		},
		{
			name: "invalid frequency",
			input: domain.Habit{
				Name:      "Read",
				Frequency: "hourly", // unsupported
			},
			wantErr:     true,
			errContains: "invalid frequency type",
		},
		{
			name: "repo error",
			input: domain.Habit{
				Name:      "Read",
				Frequency: "daily",
			},
			mockCreate: func(h *domain.Habit) error {
				return errors.New("db failed")
			},
			wantErr:     true,
			errContains: "failed to create habit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &usecase.MockHabitRepo{
				CreateFn: tt.mockCreate,
			}
			uc := &usecase.HabitUsecase{HabitRepo: mockRepo}

			err := uc.CreateHabit(&tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAllHabits(t *testing.T) {
	tests := []struct {
		name        string
		mockGetAll  func() ([]domain.Habit, error)
		wantErr     bool
		errContains string
	}{
		{
			name: "valid get all habits",
			mockGetAll: func() ([]domain.Habit, error) {
				return []domain.Habit{
					{Name: "Habit 1", Frequency: "daily"},
					{Name: "Habit 2", Frequency: "weekly"},
				}, nil
			},
			wantErr: false,
		},
		{
			name: "repo error",
			mockGetAll: func() ([]domain.Habit, error) {
				return nil, errors.New("db error")
			},
			wantErr:     true,
			errContains: "failed to get habits",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &usecase.MockHabitRepo{
				GetAllFn: tt.mockGetAll,
			}

			uc := &usecase.HabitUsecase{HabitRepo: mockRepo}
			_, err := uc.GetAllHabits()

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetHabitByID(t *testing.T) {
	tests := []struct {
		name        string
		mockGetByID func(uint) (*domain.Habit, error)
		inputID     uint
		wantErr     bool
		errContains string
		wantHabit   *domain.Habit
	}{
		{
			name:    "habit found",
			inputID: 1,
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return &domain.Habit{
					ID:        id,
					Name:      "Meditate",
					Frequency: "daily",
				}, nil
			},
			wantErr:   false,
			wantHabit: &domain.Habit{ID: 1, Name: "Meditate", Frequency: "daily"},
		},
		{
			name:    "habit not found",
			inputID: 2,
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return nil, gorm.ErrRecordNotFound
			},
			wantErr:     true,
			errContains: "habit not found",
		},
		{
			name:    "repo failure",
			inputID: 3,
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return nil, errors.New("db down")
			},
			wantErr:     true,
			errContains: "failed to retrieve habit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &usecase.MockHabitRepo{
				GetByIDFn: tt.mockGetByID,
			}
			uc := &usecase.HabitUsecase{HabitRepo: mockRepo}

			habit, err := uc.GetHabitByID(tt.inputID)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantHabit, habit)
			}
		})
	}
}

func TestUpdateHabit(t *testing.T) {
	tests := []struct {
		name        string
		inputHabit  domain.Habit
		mockGetByID func(uint) (*domain.Habit, error)
		mockUpdate  func(*domain.Habit) error
		wantErr     bool
		errContains string
	}{
		{
			name:       "successful update",
			inputHabit: domain.Habit{ID: 1, Name: "Updated Habit", Frequency: "daily"},
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return &domain.Habit{ID: id, Name: "Old Habit", Frequency: "daily"}, nil
			},
			mockUpdate: func(h *domain.Habit) error {
				return nil
			},
			wantErr: false,
		},
		{
			name:       "habit not found",
			inputHabit: domain.Habit{ID: 2, Name: "Doesn't Matter", Frequency: "daily"},
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return nil, errors.New("habit not found")
			},
			wantErr:     true,
			errContains: "failed to retrieve existing habit",
		},
		{
			name:       "empty habit name",
			inputHabit: domain.Habit{ID: 1, Name: "", Frequency: "daily"},
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return &domain.Habit{ID: id, Name: "Old", Frequency: "daily"}, nil
			},
			wantErr:     true,
			errContains: "habit name cannot be empty",
		},
		{
			name:       "invalid frequency",
			inputHabit: domain.Habit{ID: 1, Name: "Habit", Frequency: "yearly"},
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return &domain.Habit{ID: id, Name: "Old", Frequency: "daily"}, nil
			},
			wantErr:     true,
			errContains: "invalid frequency type",
		},
		{
			name:       "repository update error",
			inputHabit: domain.Habit{ID: 1, Name: "Habit", Frequency: "daily"},
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return &domain.Habit{ID: id, Name: "Old", Frequency: "daily"}, nil
			},
			mockUpdate: func(h *domain.Habit) error {
				return errors.New("db error")
			},
			wantErr:     true,
			errContains: "failed to update habit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &usecase.MockHabitRepo{
				GetByIDFn: tt.mockGetByID,
				UpdateFn:  tt.mockUpdate,
			}
			uc := &usecase.HabitUsecase{HabitRepo: mockRepo}

			err := uc.UpdateHabit(&tt.inputHabit)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteHabit(t *testing.T) {
	tests := []struct {
		name        string
		habitID     uint
		mockGetByID func(uint) (*domain.Habit, error)
		mockDelete  func(uint) error
		wantErr     bool
		errContains string
	}{
		{
			name:    "successful delete",
			habitID: 1,
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return &domain.Habit{ID: id, Name: "Test Habit"}, nil
			},
			mockDelete: func(id uint) error {
				return nil
			},
			wantErr: false,
		},
		{
			name:    "habit not found",
			habitID: 2,
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return nil, errors.New("habit not found")
			},
			wantErr:     true,
			errContains: "habit not found",
		},
		{
			name:    "delete fails",
			habitID: 3,
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return &domain.Habit{ID: id, Name: "Failing Habit"}, nil
			},
			mockDelete: func(id uint) error {
				return errors.New("db delete error")
			},
			wantErr:     true,
			errContains: "failed to delete habit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &usecase.MockHabitRepo{
				GetByIDFn: tt.mockGetByID,
				DeleteFn:  tt.mockDelete,
			}
			uc := &usecase.HabitUsecase{HabitRepo: mockRepo}

			err := uc.DeleteHabit(tt.habitID)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestMarkCompleted(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name           string
		habitID        uint
		mockGetByID    func(uint) (*domain.Habit, error)
		mockSafeUpdate func(*domain.Habit) error
		wantErr        bool
		errContains    string
		expectStreak   int
	}{
		{
			name:    "first time completion",
			habitID: 1,
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return &domain.Habit{ID: id, Frequency: "daily", LastCompletedAt: nil}, nil
			},
			mockSafeUpdate: func(h *domain.Habit) error {
				assert.Equal(t, 1, h.CurrentStreak)
				return nil
			},
			wantErr:      false,
			expectStreak: 1,
		},
		{
			name:    "daily habit continued",
			habitID: 2,
			mockGetByID: func(id uint) (*domain.Habit, error) {
				yesterday := now.Add(-23 * time.Hour)
				return &domain.Habit{ID: id, Frequency: "daily", LastCompletedAt: &yesterday, CurrentStreak: 3}, nil
			},
			mockSafeUpdate: func(h *domain.Habit) error {
				assert.Equal(t, 4, h.CurrentStreak)
				return nil
			},
			wantErr:      false,
			expectStreak: 4,
		},
		{
			name:    "weekly habit missed",
			habitID: 3,
			mockGetByID: func(id uint) (*domain.Habit, error) {
				twoWeeksAgo := now.Add(-15 * 24 * time.Hour)
				return &domain.Habit{ID: id, Frequency: "weekly", LastCompletedAt: &twoWeeksAgo, CurrentStreak: 7}, nil
			},
			mockSafeUpdate: func(h *domain.Habit) error {
				assert.Equal(t, 1, h.CurrentStreak)
				return nil
			},
			wantErr:      false,
			expectStreak: 1,
		},
		{
			name:    "habit not found",
			habitID: 4,
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return nil, errors.New("habit not found")
			},
			mockSafeUpdate: nil,
			wantErr:        true,
			errContains:    "habit not found",
		},
		{
			name:    "update fails",
			habitID: 5,
			mockGetByID: func(id uint) (*domain.Habit, error) {
				return &domain.Habit{ID: id, Frequency: "daily"}, nil
			},
			mockSafeUpdate: func(h *domain.Habit) error {
				return errors.New("db error")
			},
			wantErr:     true,
			errContains: "failed to mark habit as complete",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &usecase.MockHabitRepo{
				GetByIDFn:    tt.mockGetByID,
				SafeUpdateFn: tt.mockSafeUpdate,
			}
			uc := &usecase.HabitUsecase{HabitRepo: mockRepo}

			err := uc.MarkCompleted(tt.habitID)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetStreaks(t *testing.T) {
	tests := []struct {
		name           string
		mockGetStreaks func() ([]domain.Habit, error)
		wantErr        bool
		errContains    string
		expectedCount  int
	}{
		{
			name: "success - habits returned",
			mockGetStreaks: func() ([]domain.Habit, error) {
				return []domain.Habit{
					{ID: 1, Name: "Workout", CurrentStreak: 3},
					{ID: 2, Name: "Read", CurrentStreak: 2},
				}, nil
			},
			wantErr:       false,
			expectedCount: 2,
		},
		{
			name: "repository error",
			mockGetStreaks: func() ([]domain.Habit, error) {
				return nil, errors.New("db connection failed")
			},
			wantErr:     true,
			errContains: "failed to get all habit streaks",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &usecase.MockHabitRepo{
				GetStreaksFn: tt.mockGetStreaks,
			}
			uc := &usecase.HabitUsecase{HabitRepo: mockRepo}

			habits, err := uc.GetStreaks()

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				assert.NoError(t, err)
				assert.Len(t, habits, tt.expectedCount)
			}
		})
	}
}
