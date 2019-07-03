package payment

import (
	"time"

	"github.com/BrunoDM2943/pock_bank-api/domain"
)

func processPayments(payments []domain.Payment) error {
	for range payments {
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func processPaymentsFaster(payments []domain.Payment) error {
	for range payments {
		time.Sleep(50 * time.Millisecond)
	}
	return nil
}
