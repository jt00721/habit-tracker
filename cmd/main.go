package main

import (
	"github.com/jt00721/habit-tracker/config"
)

func main() {
	// Initialize application
	application := config.NewApp()
	// Run application
	application.Run()
}
