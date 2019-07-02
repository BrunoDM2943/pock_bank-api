package domain

import (
	"time"
)

type Account struct {
	ID		int64 `json:"id"`
	User	string `json:"name"`
	Balance	float64	`json:"balance"`
	Active	bool	`json:"active"`
	CreatedDate	time.Time `json:"created_date"`
}