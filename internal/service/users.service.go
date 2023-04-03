package service

import (
	"context"
	"errors"

	"github.com/otaxhu/bank-app/internal/entity"
	"github.com/otaxhu/bank-app/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyRegistered = errors.New("user already registered")
)

type UsersService interface {
	RegisterUser(ctx context.Context, user *entity.UserCredentials) error
	LoginUser(ctx context.Context, user *entity.UserCredentials) (*entity.DomainUser, error)
}

type usersService struct {
	userRepo repository.UsersRepository
}

func NewUsersService(userRepo repository.UsersRepository) UsersService {
	return &usersService{userRepo: userRepo}
}

func (us *usersService) RegisterUser(ctx context.Context, user *entity.UserCredentials) error {
	u, _ := us.userRepo.GetUserByEmail(ctx, user.Email)
	if u != nil {
		return ErrUserAlreadyRegistered
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	return us.userRepo.SaveUser(ctx, user)
}

func (us *usersService) LoginUser(ctx context.Context, user *entity.UserCredentials) (*entity.DomainUser, error) {
	du, err := us.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(du.Password), []byte(user.Password)); err != nil {
		return nil, err
	}

	return du, nil
}
