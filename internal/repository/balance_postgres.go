package repository

import (
	"cashflow/internal/model"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type BalancePostgres struct {
	dbPool *pgxpool.Pool
}

func NewBalancePostgres(dbPool *pgxpool.Pool) *BalancePostgres {
	return &BalancePostgres{dbPool: dbPool}
}

func (r *BalancePostgres) Deposit(input model.TransactionInput) error {
	tx, err := r.dbPool.Begin(context.Background())
	if err != nil {
		logrus.Errorf("Error starting transaction: %s", err.Error())
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
		}
	}()

	depositQuery := `
		UPDATE balances
		SET balance = balance + $1
		WHERE user_id = $2`
	res, err := tx.Exec(context.Background(), depositQuery, input.Amount, input.ToUserId)
	if err != nil {
		logrus.Errorf("Error updating balance: %s", err.Error())
		return err
	}
	if res.RowsAffected() == 0 {
		err := fmt.Errorf("balance not found for user %s", input.FromUserId)
		logrus.Errorf("balance not found for user %s", input.FromUserId)
		return err
	}

	transactionQuery := `
		INSERT INTO transactions (from_user_id, to_user_id, amount, type) 
		VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(context.Background(), transactionQuery, uuid.Nil, input.ToUserId, input.Amount, model.DepositTransactionType)
	if err != nil {
		logrus.Errorf("Error inserting transaction: %s", err.Error())
		return err
	}

	return tx.Commit(context.Background())
}

func (r *BalancePostgres) Transfer(input model.TransactionInput) (err error) {
	tx, err := r.dbPool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(context.Background())
		}
	}()

	var currentBalance float64
	checkBalanceQuery := `
		SELECT balance 
		FROM balances 
		WHERE user_id = $1`
	err = tx.QueryRow(context.Background(), checkBalanceQuery, input.FromUserId).Scan(&currentBalance)
	if err != nil {
		return err
	}
	if currentBalance < input.Amount {
		return fmt.Errorf("insufficient funds for user %s", input.FromUserId)
	}

	updatingSenderBalanceQuery := `
		UPDATE balances
		SET balance = balance - $1
		WHERE user_id = $2`
	_, err = tx.Exec(context.Background(), updatingSenderBalanceQuery, input.Amount, input.FromUserId)
	if err != nil {
		return err
	}

	updatingRecipientBalanceQuery := `
		UPDATE balances
		SET balance = balance + $1
		WHERE user_id = $2`
	res, err := tx.Exec(context.Background(), updatingRecipientBalanceQuery, input.Amount, input.ToUserId)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		err := fmt.Errorf("balance not found for user  %s", input.ToUserId)
		logrus.Errorf("balance not found for user  %s", input.ToUserId)
		return err
	}

	createTransactionQuery := `
		INSERT INTO transactions (from_user_id, to_user_id, amount, type) 
		VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(context.Background(), createTransactionQuery, input.FromUserId, input.ToUserId, input.Amount, model.TransferTransactionType)
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}
