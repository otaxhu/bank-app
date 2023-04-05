package service

import (
	"context"
	"errors"

	"github.com/otaxhu/bank-app/configs"
	"github.com/otaxhu/bank-app/internal/entity"
	"github.com/otaxhu/bank-app/internal/repository"
	"github.com/otaxhu/bank-app/internal/utils/encryption"
)

var (
	ErrUserAlreadyRegistered = errors.New("user already registered")
	ErrPasswordTooLong       = errors.New("the password is too long")
	ErrInvalidCredentials    = errors.New("the email or the password are incorrect")
)

type UsersService interface {
	RegisterUser(ctx context.Context, user *entity.UserCredentials) error
	LoginUser(ctx context.Context, user *entity.UserCredentials) (*entity.DomainUser, error)
}

type usersService struct {
	userRepo        repository.UsersRepository
	encryptionUtils *encryption.EncryptionUtils
}

func NewUsersService(cfg *configs.Configs, userRepo repository.UsersRepository) UsersService {
	return &usersService{
		userRepo:        userRepo,
		encryptionUtils: encryption.NewEncryptionUtils(cfg),
	}
}

func (us *usersService) RegisterUser(ctx context.Context, user *entity.UserCredentials) error {
	u, _ := us.userRepo.GetUserByEmail(ctx, user.Email)
	if u != nil {
		return ErrUserAlreadyRegistered
	}

	hash, err := us.encryptionUtils.GenerateHashFromPassword(user.Password)
	if err != nil {
		if err == encryption.ErrPasswordTooLong {
			return ErrPasswordTooLong
		}
		return err
	}

	user.Password = hash

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
		if err == repository.ErrResourceNotFound {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if err := us.encryptionUtils.CompareHashAndPassword(repoUser.Password, user.Password); err != nil {
		if err == encryption.ErrPasswordTooLong {
			return nil, ErrPasswordTooLong
		}
		if err == encryption.ErrMismatchedHashAndPassword {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	du := &entity.DomainUser{
		Id:    repoUser.Id,
		Roles: repoUser.Roles,
	}

	return du, nil
}
