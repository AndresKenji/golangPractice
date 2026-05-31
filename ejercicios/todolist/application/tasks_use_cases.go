package application

import (
	"errors"
	"strings"
	"time"

	"todolist/domain"
)

type LoadTasksUseCase struct {
	TaskRepo domain.TaskRepository
}

func (uc *LoadTasksUseCase) Execute() ([]*domain.Task, error) {
	return uc.TaskRepo.ListTasks()
}

type ListUserTasksUseCase struct {
	TaskRepo domain.TaskRepository
}

func (uc *ListUserTasksUseCase) Execute(userID int) ([]*domain.Task, error) {
	tasks, err := uc.TaskRepo.ListTasks()
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Task, 0)
	for _, t := range tasks {
		if t.UserID == userID {
			result = append(result, t)
		}
	}

	return result, nil
}

type CreateTaskUseCase struct {
	TaskRepo domain.TaskRepository
}

func (uc *CreateTaskUseCase) Execute(userID int, name, detail string) (*domain.Task, error) {
	name = strings.TrimSpace(name)
	detail = strings.TrimSpace(detail)
	if name == "" {
		return nil, errors.New("task name is required")
	}

	tasks, err := uc.TaskRepo.ListTasks()
	if err != nil {
		return nil, err
	}

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	newTask := domain.NewTask(name, detail)
	newTask.ID = maxID + 1
	newTask.UserID = userID
	tasks = append(tasks, newTask)

	if err = uc.TaskRepo.SaveTasks(tasks); err != nil {
		return nil, err
	}

	return newTask, nil
}

type UpdateTaskUseCase struct {
	TaskRepo domain.TaskRepository
}

func (uc *UpdateTaskUseCase) Execute(userID, taskID int, name, detail string, status domain.TaskStatus) (*domain.Task, error) {
	name = strings.TrimSpace(name)
	detail = strings.TrimSpace(detail)
	if name == "" {
		return nil, errors.New("task name is required")
	}

	if status != domain.Pending && status != domain.InProgress && status != domain.Completed {
		return nil, errors.New("invalid task status")
	}

	tasks, err := uc.TaskRepo.ListTasks()
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		if t.ID == taskID {
			if t.UserID != userID {
				return nil, errors.New("task does not belong to selected user")
			}

			t.Name = name
			t.Detail = detail
			t.Status = status
			if status == domain.Completed {
				t.FinishedAt = time.Now()
			} else {
				t.FinishedAt = time.Time{}
			}

			if err = uc.TaskRepo.UpdateTask(t); err != nil {
				return nil, err
			}

			return t, nil
		}
	}

	return nil, errors.New("task not found")
}

type DeleteTaskUseCase struct {
	TaskRepo domain.TaskRepository
}

func (uc *DeleteTaskUseCase) Execute(userID, taskID int) error {
	tasks, err := uc.TaskRepo.ListTasks()
	if err != nil {
		return err
	}

	filtered := make([]*domain.Task, 0, len(tasks))
	deleted := false
	for _, t := range tasks {
		if t.ID == taskID {
			if t.UserID != userID {
				return errors.New("task does not belong to selected user")
			}
			deleted = true
			continue
		}
		filtered = append(filtered, t)
	}

	if !deleted {
		return errors.New("task not found")
	}

	return uc.TaskRepo.SaveTasks(filtered)
}
