package payment

import (
	"testing"

	"github.com/BrunoDM2943/pock_bank-api/domain"
)

func BenchmarkProcessPayments(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processPayments(mockList(100))
	}
}

func BenchmarkProcessPaymentsFaster(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processPaymentsFaster(mockList(100))
	}
}

func BenchmarkProcessPaymentsParalel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processPaymentsParallel(mockList(100))
	}
}

func BenchmarkPayInParallel(b *testing.B) {
	processor := NewPaymentProcessor()
	user := &domain.Account{
		ID: 19,
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			processor.pay(1, 1, user)
		}
	})
}

func mockList(qtd int) []domain.Payment {
	payments := make([]domain.Payment, 0)
	for i := 0; i < qtd; i++ {
		p := domain.Payment{
			ID: i,
		}
		payments = append(payments, p)
	}
	return payments
}
