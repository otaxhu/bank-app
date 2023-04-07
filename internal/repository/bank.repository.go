package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/otaxhu/bank-app/internal/entity"
)

type BankRepository interface {
	SaveBankAccount(ctx context.Context, accountInfo *entity.BankAccount) error
	GetBankAccount(ctx context.Context, accountInfo *entity.BankAccount) (*entity.RepositoryBankAccount, error)
	UpdateBalance(ctx context.Context, repoAccount *entity.RepositoryBankAccount) error
}

type mysqlBankRepository struct {
	db *sqlx.DB
}

func NewMysqlBankRepository(db *sqlx.DB) BankRepository {
	return &mysqlBankRepository{db: db}
}
