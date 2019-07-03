package payment

import (
	"sync"
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
