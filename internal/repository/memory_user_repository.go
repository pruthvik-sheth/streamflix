package repository

import (
	"context"
	"errors"
	"streamflix/internal/domain"
	"sync"
)

type MemoryUserRepository struct {
	users map[string]*domain.User
	mu    sync.RWMutex
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

// Creates a new user to the repository
func (r *MemoryUserRepository) Create(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.Email == user.Email {
			return errors.New("User with this email already exists")
		}
	}

	r.users[user.ID] = user
	return nil
}

// Gets a user by email
func (r *MemoryUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("User not found!")
}

// Gets by id
func (r *MemoryUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("User not found!")
	}

	return user, nil
}
