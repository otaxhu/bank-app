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
	ErrPasswordTooLong       = errors.New("the password is too long")
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
		if err == bcrypt.ErrPasswordTooLong {
			return ErrPasswordTooLong
		}
		return err
	}

	user.Password = string(hash)

	if err := us.userRepo.SaveUser(ctx, user); err != nil {
		return err
	}

	u, err = us.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	return us.userRepo.SaveUserRole(ctx, u.Id, "customer")
}

func (us *usersService) LoginUser(ctx context.Context, user *entity.UserCredentials) (*entity.DomainUser, error) {
	repoUser, err := us.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(repoUser.Password), []byte(user.Password)); err != nil {
		return nil, err
	}

	du := &entity.DomainUser{
		Id:    repoUser.Id,
		Roles: repoUser.Roles,
	}

	return du, nil
}
