package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/otaxhu/bank-app/internal/entity"
)

const (
	qryInsertBankAccount = "INSERT INTO bank_accounts (id, user_id, balance, currency) VALUES (:id, :user_id, :balance, :currency)"
	qryGetBankAccount    = "SELECT id, user_id, currency, balance FROM bank_accounts WHERE user_id = ? AND currency = ?"
	qryUpdateBalanceById = "UPDATE bank_accounts SET balance = ? WHERE id = ?"
)

func (mbrepo *mysqlBankRepository) SaveBankAccount(ctx context.Context, accountInfo *entity.BankAccount) error {
	bankAccount := &entity.MysqlEntityBankAccount{
		Id:       uuid.NewString(),
		UserId:   accountInfo.UserId,
		Currency: accountInfo.Currency,
		Balance:  0,
	}
	if _, err := mbrepo.db.NamedExecContext(ctx, qryInsertBankAccount, bankAccount); err != nil {
		return err
	}
	return nil
}
func (mbrepo *mysqlBankRepository) GetBankAccount(ctx context.Context, accountInfo *entity.BankAccount) (*entity.RepositoryBankAccount, error) {
	mysqlBankAccount := &entity.MysqlEntityBankAccount{}
	if err := mbrepo.db.GetContext(ctx, mysqlBankAccount, qryGetBankAccount, accountInfo.UserId, accountInfo.Currency); err != nil {
		return nil, err
	}
	repoBankAccount := &entity.RepositoryBankAccount{
		Id:       mysqlBankAccount.Id,
		Balance:  mysqlBankAccount.Balance,
		Currency: mysqlBankAccount.Currency,
	}
	return repoBankAccount, nil
}

func (mbrepo *mysqlBankRepository) UpdateBalance(ctx context.Context, repoAccount *entity.RepositoryBankAccount) error {
	_, err := mbrepo.db.ExecContext(ctx, qryUpdateBalanceById, repoAccount.Balance, repoAccount.Id)
	if err != nil {
		return err
	}
	return nil
}
