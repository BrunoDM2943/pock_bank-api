package payment

import (
	"encoding/json"
	"os"

	"github.com/BrunoDM2943/pock_bank-api/domain"
)

func processPayments(payments []domain.Payment) error {
	file, _ := os.Create("/tmp/paymentsPointer")
	for _, payment := range payments {
		b, _ := json.Marshal(payment)
		file.Write(b)
	}
	return nil
}

func processPaymentsFaster(payments []*domain.Payment) error {
	file, _ := os.Create("/tmp/paymentsValue")
	for _, payment := range payments {
		b, _ := json.Marshal(payment)
		file.Write(b)
	}
	return nil
}
