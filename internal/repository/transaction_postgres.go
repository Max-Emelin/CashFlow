package repository

import (
	"cashflow/internal/model"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type TransactionPostgres struct {
	dbPool *pgxpool.Pool
}

func NewTransactionPostgres(dbPool *pgxpool.Pool) *TransactionPostgres {
	return &TransactionPostgres{dbPool: dbPool}
}
func (r *TransactionPostgres) GetLastTransactions(userId uuid.UUID) ([]model.Transaction, error) {
	var transactions []model.Transaction

	getTransactionsQuery := `
		SELECT id, from_user_id, to_user_id, amount, created_at, type 
		FROM transactions 
		WHERE from_user_id = :userId 
			OR to_user_id = :userId 
		ORDER BY created_at DESC 
		LIMIT 10`
	args := pgx.NamedArgs{
		"userId": userId,
	}
	rows, err := r.dbPool.Query(context.Background(), getTransactionsQuery, args)
	if err != nil {
		logrus.Errorf("Error fetching transactions for user %s: %s", userId, err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t model.Transaction
		if err := rows.Scan(&t.Id, &t.FromUserId, &t.ToUserId, &t.Amount, &t.CreatedAt, &t.Type); err != nil {
			logrus.Errorf("Error scanning transaction row: %s", err.Error())
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
