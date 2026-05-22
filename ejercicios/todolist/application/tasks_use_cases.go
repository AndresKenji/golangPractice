package application

import (
	"todolist/domain"
)

type LoadTasksUseCase struct {
	TaskRepo domain.TaskRepository
}

func (uc *LoadTasksUseCase) Execute() ([]*domain.Task, error) {
	return uc.TaskRepo.ListTasks()
}
