package infrastructure

import (
	"os"
	"encoding/json"

	"todolist/domain/task"

)

func checkFileExists(path string) (error, *os.File) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err, nil
		}
		return nil, file
	} else if err != nil {
		return err, nil
	} else {
		file, err := os.Open(path)
		if err != nil {
			return err, nil
		}
		return nil, file
	}
}

type In_memory_task_repository struct {
	Path string `default:"db.json" json:"path"`
}


func (r *In_memory_task_repository) GetTaskByName(name string) (*task.Task, error) {
	err, file := checkFileExists(r.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	var tasks []*task.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
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

func (r *In_memory_task_repository) UpdateTask(task *task.Task) error {
	err, file := checkFileExists(r.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	var tasks []*task.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			break
		}
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(task)
	if err != nil {
		return err
	}

	return nil
}

func (r *In_memory_task_repository) SaveTasks(tasks []*task.Task) error {
	err, file := checkFileExists(r.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}

func (r *In_memory_task_repository) DeleteTask(name string) error {
	err, file := checkFileExists(r.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	var tasks []*task.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.Name == name {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}

func (r *In_memory_task_repository) ListTasks() ([]*task.Task, error){
	err, file := checkFileExists(r.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []*task.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
