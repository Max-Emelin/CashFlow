package repository

import (
	"cashflow/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Balance interface {
	Deposit(input model.TransactionInput) error
	Transfer(input model.TransactionInput) error
}

type Transaction interface {
	GetLastTransactions(userId uuid.UUID) ([]model.Transaction, error)
}

type Repository struct {
	Balance
	Transaction
}

func NewRepository(dbPool *pgxpool.Pool) *Repository {
	return &Repository{
		Balance:     NewBalancePostgres(dbPool),
		Transaction: NewTransactionPostgres(dbPool),
	}
}
