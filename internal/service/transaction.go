package service

import (
	"cashflow/internal/model"
	"cashflow/internal/repository"

	uuid "github.com/google/uuid"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) GetLastTransactions(userId uuid.UUID) ([]model.Transaction, error) {
	return s.repo.GetLastTransactions(userId)
}
