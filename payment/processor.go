package payment

import (
	"sync"
	"time"

	"github.com/BrunoDM2943/pock_bank-api/domain"
)

type paymentProcessor struct {
	database map[int64][]*domain.Payment
}

func NewPaymentProcessor() paymentProcessor {
	return paymentProcessor{
		database: make(map[int64][]*domain.Payment),
	}
}

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

func processPaymentsParallel(payments []domain.Payment) error {
	wg := sync.WaitGroup{}
	for range payments {
		wg.Add(1)
		go func() {
			time.Sleep(100 * time.Millisecond)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}

func (this paymentProcessor) pay(paymentId int, value float32, user *domain.Account) {
	payment := &domain.Payment{
		ID:     paymentId,
		Value:  value,
		Date:   time.Now(),
		UserID: user.ID,
	}
	payments := append(this.database[user.ID], payment)
	this.database[user.ID] = payments

}
