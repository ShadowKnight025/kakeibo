package entity

import (

	"lionheart.dev/kakeibo-api/entity/types"
)

// Implement enum style types for Month and Category
// Implement Postgresql or Redis

type Expense struct {
	ID       string            `json:"id"`
	Title    *string           `json:"title,omitempty"`
	Category Categories        `json:"category"`
	Month    Month             `json:"month"`
	Amount   types.USD         `json:"amount"`
}
