package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jt00721/habit-tracker/internal/domain"
	"github.com/jt00721/habit-tracker/internal/usecase"
)

type HabitHandler struct {
	Usecase *usecase.HabitUsecase
}

func (handler *HabitHandler) CreateHabitApi(c *gin.Context) {
	var habit domain.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		log.Printf("Error binding json request body to create habit: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input to create habit"})
		return
	}

	err := handler.Usecase.CreateHabit(&habit)
	if err != nil {
		log.Printf("Error creating habit: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create habit. Please try again later."})
		return
	}

	c.JSON(http.StatusCreated, habit)
}

func (handler *HabitHandler) GetAllHabitsApi(c *gin.Context) {
	habits, err := handler.Usecase.GetAllHabits()
	if err != nil {
		log.Printf("Error retrieving all habits: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve all habits. Please try again later.",
		})
		return
	}

	if len(habits) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No habits found",
			"habits":  habits,
		})
		return
	}

	c.JSON(http.StatusOK, habits)
}

func (handler *HabitHandler) GetHabitByIDApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error converting habit ID URL query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid habit ID"})
		return
	}

	habit, err := handler.Usecase.GetHabitByID(uint(id))
	if err != nil {
		log.Printf("Error retrieving habit with ID(%d): %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Habit not found",
		})
		return
	}

	c.JSON(http.StatusOK, habit)
}

func (handler *HabitHandler) UpdateHabitApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error converting habit ID URL query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid habit ID"})
		return
	}

	var habit domain.Habit
	if err := c.ShouldBindJSON(&habit); err != nil {
		log.Printf("Error binding json request body to update habit: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input to update habit",
		})
		return
	}

	habit.ID = uint(id)
	err = handler.Usecase.UpdateHabit(&habit)
	if err != nil {
		log.Printf("Error updating habit with ID(%d): %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update habit. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, habit)
}

func (handler *HabitHandler) DeleteHabitApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error converting habit ID URL query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid habit ID"})
		return
	}

	err = handler.Usecase.DeleteHabit(uint(id))
	if err != nil {
		log.Printf("Error deleting habit with ID(%d): %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete habit. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Habit deleted"})
}

func (handler *HabitHandler) MarkHabitCompletedApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error converting habit ID URL query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid habit ID"})
		return
	}

	err = handler.Usecase.MarkCompleted(uint(id))
	if err != nil {
		if err.Error() == "habit not found" {
			log.Printf("Error: Tried to complete non-existing habit with ID(%d)", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
			return
		}

		log.Printf("Error marking habit with ID(%d) as complete: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark habit as completed. Please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Habit Completed"})
}

func (handler *HabitHandler) GetStreaksApi(c *gin.Context) {
	habits, err := handler.Usecase.GetStreaks()
	if err != nil {
		log.Printf("Error retrieving all habits with streaks: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve all habits with streaks. Please try again later.",
		})
		return
	}

	if len(habits) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No habits with streaks found",
			"habits":  habits,
		})
		return
	}

	c.JSON(http.StatusOK, habits)
}
