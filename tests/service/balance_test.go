package test

import (
	"cashflow/internal/model"
	"cashflow/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBalanceRepository struct {
	mock.Mock
}

func (m *MockBalanceRepository) Deposit(input model.TransactionInput) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockBalanceRepository) Transfer(input model.TransactionInput) error {
	args := m.Called(input)
	return args.Error(0)
}

func TestBalanceService_Deposit(t *testing.T) {
	mockRepo := new(MockBalanceRepository)

	service := service.NewBalanceService(mockRepo)

	t.Run("Successful deposit", func(t *testing.T) {
		mockRepo.On("Deposit", mock.AnythingOfType("model.TransactionInput")).Return(nil)

		input := model.TransactionInput{
			Amount: 100.0,
		}

		err := service.Deposit(input)
		assert.Nil(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid deposit amount", func(t *testing.T) {
		input := model.TransactionInput{
			Amount: -50.0,
		}

		err := service.Deposit(input)
		assert.NotNil(t, err)
		assert.Equal(t, "amount must be greater than zero", err.Error())
	})
}

func TestBalanceService_Transfer(t *testing.T) {
	mockRepo := new(MockBalanceRepository)

	service := service.NewBalanceService(mockRepo)

	t.Run("Successful transfer", func(t *testing.T) {
		mockRepo.On("Transfer", mock.AnythingOfType("model.TransactionInput")).Return(nil)

		input := model.TransactionInput{
			Amount: 100.0,
		}

		err := service.Transfer(input)
		assert.Nil(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid transfer amount", func(t *testing.T) {
		input := model.TransactionInput{
			Amount: -50.0,
		}

		err := service.Transfer(input)
		assert.NotNil(t, err)
		assert.Equal(t, "amount must be greater than zero", err.Error())
	})
}
