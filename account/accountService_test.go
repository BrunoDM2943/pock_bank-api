package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"errors"
	"github.com/golang/mock/gomock"
	mock "github.com/BrunoDM2943/pock_bank-api/account/mock"
	"github.com/BrunoDM2943/pock_bank-api/domain"
)

func TestSaveAccountOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDao := mock.NewMockIAccountRepo(ctrl)
	service := NewAccountService(mockDao)
	account := &domain.Account{
		User: "bdm2943",
		Balance: 100,
	}
	mockDao.EXPECT().SaveAccount(account).Return(nil).AnyTimes()
	err := service.SaveAccount(account)
	assert.Nil(t, err)
}

func TestSaveAccountNotOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDao := mock.NewMockIAccountRepo(ctrl)
	service := NewAccountService(mockDao)
	account := &domain.Account{
		User: "bdm2943",
		Balance: 100,
	}
	mockDao.EXPECT().SaveAccount(account).Return(errors.New("Error")).AnyTimes()
	err := service.SaveAccount(account)
	assert.NotNil(t, err)
}

func TestGetAccountOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDao := mock.NewMockIAccountRepo(ctrl)
	service := NewAccountService(mockDao)
	accountMock := &domain.Account{
		User: "bdm2943",
		Balance: 100,
		ID: 19,
	}
	mockDao.EXPECT().GetAccount(int64(19)).Return(accountMock, nil).AnyTimes()
	account, err := service.GetAccount(19)
	assert.Nil(t, err)
	assert.Equal(t, accountMock, account)
}

func TestGetAccountNotOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDao := mock.NewMockIAccountRepo(ctrl)
	service := NewAccountService(mockDao)
	mockDao.EXPECT().GetAccount(int64(19)).Return(nil, errors.New("Error")).AnyTimes()
	account, err := service.GetAccount(19)
	assert.Nil(t, account)
	assert.NotNil(t, err)
}