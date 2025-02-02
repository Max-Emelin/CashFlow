package service

import (
	"cashflow/internal/model"
	"cashflow/internal/repository"

	uuid "github.com/google/uuid"
)

type Balance interface {
	Deposit(input model.TransactionInput) error
	Transfer(input model.TransactionInput) error
}

type Transaction interface {
	GetLastTransactions(userId uuid.UUID) ([]model.Transaction, error)
}

type Service struct {
	Balance
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Balance:     NewBalanceService(repos.Balance),
		Transaction: NewTransactionService(repos.Transaction),
	}
}
