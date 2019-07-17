package account

import "github.com/BrunoDM2943/pock_bank-api/domain"

//Service for account services
//go:generate mockgen -source=./accountService.go -destination=./mock/mock_accountService.go
type Service interface {
	SaveAccount(person *domain.Account) error
	GetAccount(ID int64) (*domain.Account, error)
}

func NewAccountService(repo Repository) Service {
	return &accountService{repo}
}

type accountService struct {
	repo Repository
}

func (service accountService) SaveAccount(account *domain.Account) error {
	return service.repo.SaveAccount(account)
}

func (service accountService) GetAccount(ID int64) (*domain.Account, error) {
	return service.repo.GetAccount(ID)
}
