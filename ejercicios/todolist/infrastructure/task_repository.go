package infrastructure

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"todolist/domain"
)

type In_memory_task_repository struct {
	Path string `default:"db.json" json:"path"`
}

func (r *In_memory_task_repository) dbPath() string {
	if strings.TrimSpace(r.Path) == "" {
		return "db.json"
	}

	return r.Path
}

func (r *In_memory_task_repository) ensureDBFile() error {
	path := r.dbPath()
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.WriteFile(path, []byte("[]"), 0644)
	}

	return err
}

func (r *In_memory_task_repository) loadTasks() ([]*domain.Task, error) {
	err := r.ensureDBFile()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(r.dbPath())
	if err != nil {
		return nil, err
	}

	if len(strings.TrimSpace(string(data))) == 0 {
		return []*domain.Task{}, nil
	}

	var tasks []*domain.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	if tasks == nil {
		return []*domain.Task{}, nil
	}

	return tasks, nil
}

func (r *In_memory_task_repository) saveTasks(tasks []*domain.Task) error {
	err := r.ensureDBFile()
	if err != nil {
		return err
	}

	if tasks == nil {
		tasks = []*domain.Task{}
	}

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.dbPath(), data, 0644)
}

func (r *In_memory_task_repository) GetTaskByName(name string) (*domain.Task, error) {
	tasks, err := r.loadTasks()
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		if t.Name == name {
			return t, nil
		}
	}

	return nil, nil
}

func (r *In_memory_task_repository) UpdateTask(taskToUpdate *domain.Task) error {
	if taskToUpdate == nil {
		return errors.New("task is nil")
	}

	tasks, err := r.loadTasks()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == taskToUpdate.ID {
			tasks[i] = taskToUpdate
			found = true
			break
		}
	}

	if !found {
		return errors.New("task not found")
	}

	return r.saveTasks(tasks)
}

func (r *In_memory_task_repository) SaveTasks(tasks []*domain.Task) error {
	return r.saveTasks(tasks)
}

func (r *In_memory_task_repository) DeleteTask(name string) error {
	tasks, err := r.loadTasks()
	if err != nil {
		return err
	}

	originalLen := len(tasks)
	for i, t := range tasks {
		if t.Name == name {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	if len(tasks) == originalLen {
		return errors.New("task not found")
	}

	return r.saveTasks(tasks)
}

func (r *In_memory_task_repository) ListTasks() ([]*domain.Task, error) {
	return r.loadTasks()
}
