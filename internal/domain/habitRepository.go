package domain

type HabitRepository interface {
	Create(h *Habit) error
	GetByID(id uint) (*Habit, error)
	GetAll() ([]Habit, error)
	Update(h *Habit) error
	Delete(id uint) error
	SafeUpdate(h *Habit) error
	GetStreaks() ([]Habit, error)
}
