package domain

import (
	"github.com/Banking-System/dto"
	"github.com/Banking-System/errs"
)

const dbTSLayout = "2022-03-30 20:00:00"

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
	return &dto.NewAccountResponse{AccountId: a.AccountId} //  初始化,安全问题
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	//在这里添加方法
	SavaTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}

//
func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount { //static check
		return false
	}
	return true

	// return !a.Amount >= amount
}

//
func NewAccount(customerId, accountType string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: dbTSLayout,
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}
