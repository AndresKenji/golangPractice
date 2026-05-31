package infrastructure

import (
	"errors"
	"os"
	"strings"

	"encoding/json"

	"todolist/domain"
)

type In_memory_user_repository struct {
	Path string `default:"users.json" json:"path"`
}

func (r *In_memory_user_repository) dbPath() string {
	if strings.TrimSpace(r.Path) == "" {
		return "users.json"
	}

	return r.Path
}

func (r *In_memory_user_repository) ensureDBFile() error {
	path := r.dbPath()
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.WriteFile(path, []byte("[]"), 0644)
	}

	return err
}

func (r *In_memory_user_repository) loadUsers() ([]*domain.User, error) {
	err := r.ensureDBFile()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(r.dbPath())
	if err != nil {
		return nil, err
	}

	if len(strings.TrimSpace(string(data))) == 0 {
		return []*domain.User{}, nil
	}

	var users []*domain.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	if users == nil {
		return []*domain.User{}, nil
	}

	return users, nil
}

func (r *In_memory_user_repository) saveUsers(users []*domain.User) error {
	err := r.ensureDBFile()
	if err != nil {
		return err
	}

	if users == nil {
		users = []*domain.User{}
	}

	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.dbPath(), data, 0644)
}

func (r *In_memory_user_repository) CreateUser(user *domain.User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	users, err := r.loadUsers()
	if err != nil {
		return err
	}

	for _, u := range users {
		if strings.EqualFold(u.Name, user.Name) {
			return errors.New("user already exists")
		}
	}

	users = append(users, user)
	return r.saveUsers(users)
}

func (r *In_memory_user_repository) GetUserByName(name string) (*domain.User, error) {
	users, err := r.loadUsers()
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		if strings.EqualFold(u.Name, name) {
			return u, nil
		}
	}

	return nil, nil
}

func (r *In_memory_user_repository) UpdateUser(user *domain.User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	users, err := r.loadUsers()
	if err != nil {
		return err
	}

	found := false
	for i, u := range users {
		if u.ID == user.ID {
			users[i] = user
			found = true
			break
		}
	}

	if !found {
		return errors.New("user not found")
	}

	return r.saveUsers(users)
}

func (r *In_memory_user_repository) DeleteUser(name string) error {
	users, err := r.loadUsers()
	if err != nil {
		return err
	}

	originalLen := len(users)
	for i, u := range users {
		if strings.EqualFold(u.Name, name) {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	if len(users) == originalLen {
		return errors.New("user not found")
	}

	return r.saveUsers(users)
}

func (r *In_memory_user_repository) ListUsers() ([]*domain.User, error) {
	return r.loadUsers()
}
