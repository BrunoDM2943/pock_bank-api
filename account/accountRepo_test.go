package account

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/BrunoDM2943/pock_bank-api/domain"
	"github.com/stretchr/testify/assert"
)

func TestSaveAccountOnDBWihoutID(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "../gotest.db")
	db.DropTable(&domain.Account{})
	db.AutoMigrate(&domain.Account{})
	repo := NewAccountRepo(db)
	defer db.Close()
	account := &domain.Account{
		User: "Test",
	}
	assert.Nil(t, repo.SaveAccount(account))
	assert.NotEqual(t, 0, account.ID)
}

func TestSaveAccountOnDBWithID(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "../gotest.db")
	db.DropTable(&domain.Account{})
	db.AutoMigrate(&domain.Account{})
	repo := NewAccountRepo(db)
	defer db.Close()
	account := &domain.Account{
		User: "Test",
		ID:   19,
	}
	assert.Nil(t, repo.SaveAccount(account))
	assert.Equal(t, int64(19), account.ID)
}

func TestSaveAccountOnDBError(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "../gotest.db")
	db.DropTable(&domain.Account{})
	repo := NewAccountRepo(db)
	defer db.Close()
	account := &domain.Account{
		User: "Test",
		ID:   19,
	}
	assert.NotNil(t, repo.SaveAccount(account))
}

func TestGetAccount(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "../gotest.db")
	db.DropTable(&domain.Account{})
	db.AutoMigrate(&domain.Account{})
	repo := NewAccountRepo(db)
	defer db.Close()
	account := &domain.Account{
		User: "Test",
		ID:   19,
	}
	repo.SaveAccount(account)
	a, _ := repo.GetAccount(19)
	assert.Equal(t, account.User, a.User)
}

func TestGetAccountError(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "../gotest.db")
	db.DropTable(&domain.Account{})
	repo := NewAccountRepo(db)
	defer db.Close()
	_, err := repo.GetAccount(19)
	assert.NotNil(t, err)
}