package repository

import (
	"cashflow/internal/model"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
	defer tx.Rollback(context.Background())

	depositQuery := `
		UPDATE balances
		SET balance = balance + :amount
		WHERE user_id = :userId`
	args := pgx.NamedArgs{
		"amount": input.Amount,
		"userId": input.ToUserId,
	}
	_, err = tx.Exec(context.Background(), depositQuery, args)
	if err != nil {
		logrus.Errorf("Error updating balance: %s", err.Error())
		return err
	}

	transactionQuery := `
		INSERT INTO transactions (from_user_id, to_user_id, amount, type) 
		VALUES (:from_user_id, :to_user_id, :amount, :type)`
	args = pgx.NamedArgs{
		"from_user_id": uuid.Nil,
		"to_user_id":   input.ToUserId,
		"amount":       input.Amount,
		"type":         model.DepositTransactionType,
	}
	_, err = tx.Exec(context.Background(), transactionQuery, args)
	if err != nil {
		logrus.Errorf("Error inserting transaction: %s", err.Error())
		return err
	}

	return tx.Commit(context.Background())
}

func (r *BalancePostgres) Transfer(input model.TransactionInput) error {
	tx, err := r.dbPool.Begin(context.Background())
	if err != nil {
		logrus.Errorf("Error starting transaction: %s", err.Error())
		return err
	}
	defer tx.Rollback(context.Background())

	updatingSenderBalanceQuery := `
		UPDATE balances
		SET balance = balance - :amount
		WHERE user_id = :userId
			AND balance >= :amount`
	args := pgx.NamedArgs{
		"amount": input.Amount,
		"userId": input.FromUserId,
	}
	_, err = tx.Exec(context.Background(), updatingSenderBalanceQuery, args)
	if err != nil {
		logrus.Errorf("Error updating sender balance: %s", err.Error())
		return err
	}

	updatingRecipientBalanceQuery := `
		UPDATE balances
		SET balance = balance + :amount
		WHERE user_id = :userId`
	args = pgx.NamedArgs{
		"amount": input.Amount,
		"userId": input.ToUserId,
	}
	_, err = tx.Exec(context.Background(), updatingRecipientBalanceQuery, args)
	if err != nil {
		logrus.Errorf("Error updating recipient balance: %s", err.Error())
		return err
	}

	createTransactionQuery := `
		INSERT INTO transactions (from_user_id, to_user_id, amount, type) 
		VALUES (:from_user_id, :to_user_id, :amount, :type)`
	args = pgx.NamedArgs{
		"from_user_id": input.FromUserId,
		"to_user_id":   input.ToUserId,
		"amount":       input.Amount,
		"type":         model.TransferTransactionType,
	}
	_, err = tx.Exec(context.Background(), createTransactionQuery, args)
	if err != nil {
		logrus.Errorf("Error creating transaction: %s", err.Error())
		return err
	}

	return tx.Commit(context.Background())
}
