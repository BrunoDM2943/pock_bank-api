package account

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mock "github.com/BrunoDM2943/pock_bank-api/account/mock"
	"github.com/BrunoDM2943/pock_bank-api/domain"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPostAccountOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock.NewMockService(ctrl)
	personHandler := NewAccountHandler(
		service,
	)

	service.EXPECT().SaveAccount(gomock.Any()).DoAndReturn(func(account *domain.Account) error {
		account.ID = 1
		return nil
	})

	req, _ := http.NewRequest("POST", "/api/account", bytes.NewBufferString(`
		{
			"user": "bdm2943"
		}
	`))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(personHandler.PostAccount)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestPostAccountNOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mock.NewMockRepository(ctrl)
	personHandler := NewAccountHandler(
		service,
	)

	service.EXPECT().SaveAccount(gomock.Any()).Return(errors.New(""))

	req, _ := http.NewRequest("POST", "/api/account", bytes.NewBufferString(`
		{
			"user": "bdm2943"
		}
	`))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(personHandler.PostAccount)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
