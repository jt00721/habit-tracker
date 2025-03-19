package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jt00721/habit-tracker/internal/handler"
	"github.com/jt00721/habit-tracker/internal/usecase"
)

func SetupRoutes(router *gin.Engine, uc *usecase.HabitUsecase) {
	habitHandler := &handler.HabitHandler{Usecase: uc}

	router.POST("/api/habits", habitHandler.CreateHabitApi)
	router.GET("/api/habits", habitHandler.GetAllHabitsApi)
	router.GET("/api/habits/:id", habitHandler.GetHabitByIDApi)
	router.PUT("/api/habits/:id", habitHandler.UpdateHabitApi)
	router.DELETE("/api/habits/:id", habitHandler.DeleteHabitApi)
	router.GET("/api/habits/with_streaks", habitHandler.GetStreaksApi)
	router.PATCH("/api/habits/:id/mark_complete", habitHandler.MarkHabitCompletedApi)
}
