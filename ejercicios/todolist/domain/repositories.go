package domain

type TaskRepository interface  {
	CreateTask(task *Task) error
	GetTaskByName(name string) (*Task, error)
	UpdateTask(task *Task) error
	SaveTask(task *Task) error
	DeleteTask(name string) error
	ListTasks() ([]*Task, error)
}

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByName(name string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(name string) error
	ListUsers() ([]*User, error)
}