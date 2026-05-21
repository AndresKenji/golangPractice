package domain

import "time"

type Task struct {
	ID         int        `json:"id"`
	UserID     int        `json:"user_id"`
	Name       string     `json:"name"`
	Detail     string     `json:"detail"`
	CreatedAt  time.Time  `json:"created_at"`
	FinishedAt time.Time  `json:"finished_at"`
	Status     TaskStatus `json:"status"`
}

func NewTask(name, detail string) *Task {
	return &Task{
		Name:      name,
		Detail:    detail,
		CreatedAt: time.Now(),
		Status:    Pending,
	}
}

func (t *Task) Complete() {
	if t.Status == Completed {
		return
	}
	t.Status = Completed
	t.FinishedAt = time.Now()
}
