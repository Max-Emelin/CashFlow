package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	DepositTransactionType  = "deposit"
	TransferTransactionType = "transfer"
)

type Transaction struct {
	Id         uuid.UUID `json:"id" db:"id"`
	FromUserId uuid.UUID `json:"from_user_id" db:"from_user_id"`
	ToUserId   uuid.UUID `json:"to_user_id" db:"to_user_id" binding:"required"`
	Amount     float64   `json:"amount" db:"amount" binding:"required"`
	CreatedAt  time.Time `json:"created_at" db:"created_at" binding:"required"`
	Type       string    `json:"type" db:"type" binding:"required"`
}

type TransactionInput struct {
	FromUserId *uuid.UUID `json:"from_user_id"`
	ToUserId   uuid.UUID  `json:"to_user_id"`
	Amount     float64    `json:"amount"`
}
