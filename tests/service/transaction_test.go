package test

import (
	"cashflow/internal/model"
	"cashflow/internal/service"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) GetLastTransactions(userId uuid.UUID) ([]model.Transaction, error) {
	args := m.Called(userId)
	return args.Get(0).([]model.Transaction), args.Error(1)
}

func TestTransactionService_GetLastTransactions(t *testing.T) {
	mockRepo := new(MockTransactionRepository)

	service := service.NewTransactionService(mockRepo)

	t.Run("Successful get last transactions", func(t *testing.T) {
		userId := uuid.New()
		mockRepo.On("GetLastTransactions", userId).Return([]model.Transaction{
			{FromUserId: userId, ToUserId: uuid.New(), Amount: 100.0},
		}, nil)

		transactions, err := service.GetLastTransactions(userId)
		assert.Nil(t, err)
		assert.Len(t, transactions, 1)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Nil when getting last transactions", func(t *testing.T) {
		userId := uuid.New()
		mockRepo.On("GetLastTransactions", userId).Return([]model.Transaction{}, nil)

		transactions, err := service.GetLastTransactions(userId)

		assert.Nil(t, err)
		assert.Empty(t, transactions)
		mockRepo.AssertExpectations(t)
	})
}
