package infrastructure

import (
	"os"
	"todolist/domain/task"
)

func validateFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

type In_memory_task_repository struct {
	Path string `default:"db.json" json:"path"`
}

func (r *In_memory_task_repository)CreateTask(task *task.Task) error {
	if err := validateFile(r.Path); err != nil {
		return err
	}


	return nil
}

func (r *In_memory_task_repository)GetTaskByName(name string) (*task.Task, error)

func (r *In_memory_task_repository)UpdateTask(task *task.Task) error

func (r *In_memory_task_repository)SaveTask(task *task.Task) error

func (r *In_memory_task_repository)DeleteTask(name string) error

func (r *In_memory_task_repository)ListTasks() ([]*task.Task, error)