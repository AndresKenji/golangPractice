package application

import (
	"errors"
	"strings"
	"time"

	"todolist/domain"
)

type ListUsersUseCase struct {
	UserRepo domain.UserRepository
}

func (uc *ListUsersUseCase) Execute() ([]*domain.User, error) {
	return uc.UserRepo.ListUsers()
}

type SelectOrCreateUserUseCase struct {
	UserRepo domain.UserRepository
}

func (uc *SelectOrCreateUserUseCase) Execute(name string) (*domain.User, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("user name is required")
	}

	existing, err := uc.UserRepo.GetUserByName(name)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return existing, nil
	}

	users, err := uc.UserRepo.ListUsers()
	if err != nil {
		return nil, err
	}

	maxID := 0
	for _, u := range users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}

	user := &domain.User{
		ID:        maxID + 1,
		Name:      name,
		CreatedAt: time.Now().Format(time.RFC3339),
		Active:    true,
	}

	if err = uc.UserRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
