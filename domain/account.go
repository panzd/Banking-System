package domain

import (
	"github.com/Banking-System/dto"
	"github.com/Banking-System/errs"
)

type Account struct {
	AccountId   string
	Customerid  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
