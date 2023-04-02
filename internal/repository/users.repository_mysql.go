package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/otaxhu/bank-app/internal/entity"
)

const (
	qryInsertUser     = "INSERT INTO users (id, email, password) VALUES (:id, :email, :password)"
	qryGetUserByEmail = "SELECT id, email, password FROM users WHERE email = ?"
)

func (murepo *mysqlUsersRepo) SaveUser(ctx context.Context, user *entity.UserCredentials) error {
	mysqlUser := &entity.MysqlEntityUser{
		Id:       uuid.NewString(),
		Email:    user.Email,
		Password: user.Password,
	}
	_, err := murepo.db.NamedExecContext(ctx, qryInsertUser, mysqlUser)
	if err != nil {
		return err
	}
	return nil
}

func (murepo *mysqlUsersRepo) GetUserByEmail(ctx context.Context, email string) (*entity.DomainUser, error) {
	mysqlUser := &entity.MysqlEntityUser{}
	if err := murepo.db.GetContext(ctx, mysqlUser, qryGetUserByEmail, email); err != nil {
		return nil, err
	}
	return &entity.DomainUser{
		Id:       mysqlUser.Id,
		Email:    mysqlUser.Email,
		Password: mysqlUser.Password,
	}, nil
}
