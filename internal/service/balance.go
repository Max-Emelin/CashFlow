package service

import (
	"cashflow/internal/model"
	"cashflow/internal/repository"
)

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) Deposit(input model.TransactionInput) error {
	return s.repo.Deposit(input)
}

func (s *BalanceService) Transfer(input model.TransactionInput) error {
	return s.repo.Transfer(input)
}
