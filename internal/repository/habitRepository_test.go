package repository_test

import (
	"testing"

	"github.com/jt00721/habit-tracker/internal/domain"
	"github.com/jt00721/habit-tracker/internal/repository"
	testutils "github.com/jt00721/habit-tracker/test"
	"github.com/stretchr/testify/assert"
)

func TestGetHabitByID(t *testing.T) {
	// Set up test DB
	db, teardown := testutils.NewTestDB(t)
	defer teardown()

	repo := &repository.HabitRepository{DB: db}

	// Retrieve habit by ID
	habit, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Test Habit", habit.Name)
	assert.Equal(t, "daily", habit.Frequency)
	assert.Equal(t, 5, habit.CurrentStreak)
}

func TestCreate(t *testing.T) {
	// Set up test DB
	db, teardown := testutils.NewTestDB(t)
	defer teardown()

	repo := &repository.HabitRepository{DB: db}

	habit := domain.Habit{
		Name:      "Test Create Habit",
		Frequency: "weekly",
	}

	err := repo.Create(&habit)

	assert.NoError(t, err)

	// Fetch created habit
	createdHabit, err := repo.GetByID(habit.ID)
	assert.NoError(t, err)

	assert.Equal(t, 0, createdHabit.CurrentStreak)
	assert.Nil(t, createdHabit.LastCompletedAt)
	assert.Equal(t, 0, createdHabit.TotalCompletions)
}

func TestGetAll(t *testing.T) {
	db, teardown := testutils.NewTestDB(t)
	defer teardown()

	repo := &repository.HabitRepository{DB: db}

	repo.Create(&domain.Habit{
		Name:      "Test Habit 2",
		Frequency: "weekly",
	})

	repo.Create(&domain.Habit{
		Name:      "Test Habit 3",
		Frequency: "daily",
	})

	// Retrieve habit by ID
	habits, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, habits, 3)
}

func TestUpdate(t *testing.T) {
	// Set up test DB
	db, teardown := testutils.NewTestDB(t)
	defer teardown()

	repo := &repository.HabitRepository{DB: db}

	existingHabit, err := repo.GetByID(1)
	assert.NoError(t, err)

	existingHabit.Name = "Updated Test Habit"
	existingHabit.Frequency = "weekly"
	err = repo.Update(existingHabit)
	assert.NoError(t, err)

	// Fetch created habit
	updatedHabit, err := repo.GetByID(existingHabit.ID)
	assert.NoError(t, err)

	assert.Equal(t, "Updated Test Habit", updatedHabit.Name)
	assert.Equal(t, "weekly", updatedHabit.Frequency)
}

func TestDelete(t *testing.T) {
	// Set up test DB
	db, teardown := testutils.NewTestDB(t)
	defer teardown()

	repo := &repository.HabitRepository{DB: db}

	_, err := repo.GetByID(1)
	assert.NoError(t, err)

	// Retrieve habit by ID
	err = repo.Delete(1)
	assert.NoError(t, err)

	_, err = repo.GetByID(1)
	assert.Error(t, err)
}
