package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/otaxhu/bank-app/internal/entity"
)

type UsersRepository interface {
	SaveUser(ctx context.Context, user *entity.UserCredentials) error
	GetUserByEmail(ctx context.Context, email string) (*entity.DomainUser, error)
}

type mysqlUsersRepo struct {
	db *sqlx.DB
}

func NewMysqlUsersRepository(db *sqlx.DB) UsersRepository {
	return &mysqlUsersRepo{db: db}
}
