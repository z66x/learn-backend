package main
import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	StatusTodo       = "to-do"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

func NewTask(id int, title string) *Task {
	now := time.Now()
	return &Task{
		ID:        id,
		Title:     title,
		Status:    StatusTodo,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
