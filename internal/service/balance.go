package service

import (
	"cashflow/internal/model"
	"cashflow/internal/repository"
	"errors"
)

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) Deposit(input model.TransactionInput) error {
	if input.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return s.repo.Deposit(input)
}

func (s *BalanceService) Transfer(input model.TransactionInput) error {
	if input.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return s.repo.Transfer(input)
}
