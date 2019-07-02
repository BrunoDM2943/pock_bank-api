package account

import (
	"github.com/jinzhu/gorm"
	"github.com/BrunoDM2943/pock_bank-api/domain"
)

//IAccountRepo repositoyry
//go:generate mockgen -source=./accountRepo.go -destination=./mock/mock_accountRepo.go
type IAccountRepo interface {
	SaveAccount(account *domain.Account) error
	GetAccount(ID int64) (*domain.Account, error)
}

//NewAccountRepo creates a IAccountRepo object
func NewAccountRepo(db *gorm.DB) IAccountRepo {
	return &accountRepo{
		db,
	}
}

type accountRepo struct {
	db *gorm.DB
}

func (repo accountRepo) SaveAccount(Account *domain.Account) error {
	if Account.ID == 0 {
		Account.ID = domain.GenerateID()
	}
	return repo.db.Create(Account).Error
}

func (repo accountRepo) GetAccount(ID int64) (*domain.Account, error) {
	Account := &domain.Account{}
	err := repo.db.Find(Account, ID).Error
	return Account, err
}