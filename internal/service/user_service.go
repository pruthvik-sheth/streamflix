package service

import (
	"context"
	"errors"
	"streamflix/internal/domain"
	"streamflix/internal/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Register(ctx context.Context, email string, password string) (*domain.User, error) {
	// Validate input
	if email == "" || password == "" {
		return nil, errors.New("Email and password are required!")
	}

	if len(password) < 8 {
		return nil, errors.New("Password must be 8 characters long!")
	}

	existingUser, _ := s.userRepo.GetByEmail(ctx, email)

	if existingUser != nil {
		return nil, errors.New("User with this email already exists!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	// Create a user
	user := &domain.User{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.userRepo.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *UserService) Login(ctx context.Context, email string, password string) (*domain.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("Email and password are required!")
	}

	// Check if the user exists
	existingUser, err := s.userRepo.GetByEmail(ctx, email)

	if err != nil {
		return nil, errors.New("Invalid Credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid Credentials")
	}

	// If the user exists then the next step is to check whether the user has entered the correct password
	return existingUser, nil
}
