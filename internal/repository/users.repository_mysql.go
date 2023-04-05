package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/otaxhu/bank-app/internal/entity"
)

const (
	qryInsertUser       = "INSERT INTO users (id, email, password) VALUES (:id, :email, :password)"
	qryInsertUserRole   = "INSERT INTO user_roles (id, role, user_id) VALUES (:id, :role, :user_id)"
	qryGetUserByEmail   = "SELECT id, email, password FROM users WHERE email = ?"
	qryGetUserById      = "SELECT id, email, password FROM users WHERE id = ?"
	qryGetUserRolesById = "SELECT id, role, user_id FROM user_roles WHERE user_id = ?"
	qryDeleteUser       = "DELETE FROM users WHERE id = ?"
	qryUpdateUser       = "UPDATE users SET email = :email, password = :password WHERE id = :id"
)

var (
	errNoRowsAffected = errors.New("no rows affected in the query sql")
)

func (murepo *mysqlUsersRepo) SaveUser(ctx context.Context, user *entity.UserCredentials) error {
	mysqlUser := &entity.MysqlEntityUser{
		Id:       uuid.NewString(),
		Email:    user.Email,
		Password: user.Password,
	}
	if _, err := murepo.db.NamedExecContext(ctx, qryInsertUser, mysqlUser); err != nil {
		return err
	}
	return nil
}

func (murepo *mysqlUsersRepo) GetUserByEmail(ctx context.Context, email string) (*entity.RepositoryUser, error) {
	mysqlUser := &entity.MysqlEntityUser{}
	if err := murepo.db.GetContext(ctx, mysqlUser, qryGetUserByEmail, email); err != nil {
		return nil, err
	}
	mysqlUserRoles := []entity.MysqlEntityUserRole{}
	if err := murepo.db.SelectContext(ctx, &mysqlUserRoles, qryGetUserRolesById, mysqlUser.Id); err != nil {
		return nil, err
	}
	repoUser := &entity.RepositoryUser{
		Id:       mysqlUser.Id,
		Email:    mysqlUser.Email,
		Password: mysqlUser.Password,
	}
	for _, ur := range mysqlUserRoles {
		repoUser.Roles = append(repoUser.Roles, ur.Role)
	}
	return repoUser, nil
}

func (murepo *mysqlUsersRepo) GetUserById(ctx context.Context, userId string) (*entity.RepositoryUser, error) {
	mysqlUser := &entity.MysqlEntityUser{}
	if err := murepo.db.GetContext(ctx, mysqlUser, qryGetUserById, userId); err != nil {
		return nil, err
	}
	mysqlUserRoles := []entity.MysqlEntityUserRole{}
	if err := murepo.db.SelectContext(ctx, &mysqlUserRoles, qryGetUserRolesById, userId); err != nil {
		return nil, err
	}
	repoUser := &entity.RepositoryUser{
		Id:       mysqlUser.Id,
		Email:    mysqlUser.Email,
		Password: mysqlUser.Password,
	}
	for _, ur := range mysqlUserRoles {
		repoUser.Roles = append(repoUser.Roles, ur.Role)
	}
	return repoUser, nil
}

func (murepo *mysqlUsersRepo) DeleteUser(ctx context.Context, userId string) error {
	result, err := murepo.db.ExecContext(ctx, qryDeleteUser, userId)
	if rowsAffected, _ := result.RowsAffected(); rowsAffected <= 0 {
		return errNoRowsAffected
	}
	if err != nil {
		return err
	}
	return nil
}

func (murepo *mysqlUsersRepo) UpdateUser(ctx context.Context, user *entity.RepositoryUser) error {
	mysqlUser := &entity.MysqlEntityUser{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
	}
	result, err := murepo.db.NamedExecContext(ctx, qryUpdateUser, mysqlUser)
	if err != nil {
		return err
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected <= 0 {
		return errNoRowsAffected
	}
	return nil
}

func (murepo *mysqlUsersRepo) SaveUserRole(ctx context.Context, userId, role string) error {
	mysqlUserRole := &entity.MysqlEntityUserRole{
		Id:     uuid.NewString(),
		Role:   role,
		UserId: userId,
	}
	if _, err := murepo.db.NamedExecContext(ctx, qryInsertUserRole, mysqlUserRole); err != nil {
		return err
	}
	return nil
}
