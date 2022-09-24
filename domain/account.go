package domain

import (
	"github.com/Banking-System/dto"
	"github.com/Banking-System/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	//return &dto.NewAccountResponse{a.AccountId}
	// Go vet的检查，需要初始化
	// 否则会返回无键字段的错误 literal uses unkeyed fields
	return &dto.NewAccountResponse{AccountId: a.AccountId} //  初始化
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
