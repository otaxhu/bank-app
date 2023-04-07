package service

import (
	"context"
	"errors"

	"github.com/otaxhu/bank-app/internal/entity"
	"github.com/otaxhu/bank-app/internal/repository"
)

var (
	ErrBankAccountAlreadyExists = errors.New("already exists a bank account with the currency provided")
	ErrInsufficientFunds        = errors.New("insufficient funds")
)

type BankService interface {
	CreateBankAccount(ctx context.Context, accountInfo *entity.BankAccount) error
	SendMoneyFromTo(ctx context.Context, userAccountInfo, destUserAccountInfo *entity.BankAccount, amount float64) error
	CheckBankAccount(ctx context.Context, accountInfo *entity.BankAccount) (*entity.DomainBankAccount, error)
}

type bankService struct {
	bankRepo repository.BankRepository
	userRepo repository.UsersRepository
}

func NewBankService(br repository.BankRepository, ur repository.UsersRepository) BankService {
	return &bankService{
		bankRepo: br,
		userRepo: ur,
	}
}

func (bs *bankService) CreateBankAccount(ctx context.Context, accountInfo *entity.BankAccount) error {
	account, _ := bs.bankRepo.GetBankAccount(ctx, accountInfo)
	if account != nil {
		return ErrBankAccountAlreadyExists
	}
	return bs.bankRepo.SaveBankAccount(ctx, accountInfo)
}

func (bs *bankService) SendMoneyFromTo(ctx context.Context, userAccountInfo, destUserAccount *entity.BankAccount, amount float64) error {
	repoBankAccount, err := bs.bankRepo.GetBankAccount(ctx, userAccountInfo)
	if err != nil {
		return err
	}
	repoDestUser, err := bs.bankRepo.GetBankAccount(ctx, destUserAccount)
	if err != nil {
		return err
	}
	var tax float64
	if repoBankAccount.Currency != repoDestUser.Currency {

		// Impuesto por la conversion
		tax = amount * 0.1

		// Conversion de dolares a bolivares
		if repoBankAccount.Currency == "USD" && repoDestUser.Currency == "VES" {
			repoDestUser.Balance += amount * 25
		}
		// Conversion de bolivares a dolares
		if repoBankAccount.Currency == "VES" && repoDestUser.Currency == "USD" {
			repoDestUser.Balance += amount / 25
		}
	} else {
		repoDestUser.Balance += amount
	}
	repoBankAccount.Balance -= (tax + amount)
	if repoBankAccount.Balance < 0 {
		return ErrInsufficientFunds
	}

	if err := bs.bankRepo.UpdateBalance(ctx, repoBankAccount); err != nil {
		return err
	}
	return bs.bankRepo.UpdateBalance(ctx, repoDestUser)
}

func (bs *bankService) CheckBankAccount(ctx context.Context, accountInfo *entity.BankAccount) (*entity.DomainBankAccount, error) {
	repoBankAccount, err := bs.bankRepo.GetBankAccount(ctx, accountInfo)
	if err != nil {
		return nil, err
	}
	domainBankAccount := &entity.DomainBankAccount{
		Balance:  repoBankAccount.Balance,
		Currency: repoBankAccount.Currency,
	}
	return domainBankAccount, nil
}
