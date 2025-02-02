package model

import "github.com/google/uuid"

type Balance struct {
	Id      uuid.UUID `json:"id" db:"id"`
	UserId  uuid.UUID `json:"user_id" db:"user_id" binding:"required"`
	Balance float64   `json:"balance" db:"balance" binding:"required"`
}
