package account

import "github.com/BrunoDM2943/pock_bank-api/domain"

type IAccountService interface {
	SaveAccount(person *domain.Account) error
	GetAccount(ID int64) (*domain.Account, error)
}

func NewAccountService(repo IAccountRepo) IAccountService {
	return &accountService{repo}
}

type accountService struct {
	repo IAccountRepo
}

func (service accountService) SaveAccount(person *domain.Account) error {
	return service.repo.SaveAccount(person)
}

func (service accountService) GetAccount(ID int64) (*domain.Account, error) {
	return service.repo.GetAccount(ID)
}