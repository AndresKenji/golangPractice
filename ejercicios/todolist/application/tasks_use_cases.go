package application

import (
	"todolist/domain/task"
)


type LoadTasksUseCase struct {
	TaskRepo task.TaskRepository
}

func (uc *LoadTasksUseCase) Execute() (*task.Task, error) {
	return uc.TaskRepo.ListTasks()
}
