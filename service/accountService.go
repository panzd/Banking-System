package service

import (
	"time"

	"github.com/Banking-System/domain"
	"github.com/Banking-System/dto"
	"github.com/Banking-System/errs"
)

const dbTSLayout = "2022-03-30 20:00:00"

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) // 要在接口中实现方法
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2022-03-30 20:00:00"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	// account := domain.NewAccount(req.CustomerId,req.AccountType,req.Amount)

	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResponseDto() // 返回值是引用

	return response, nil //已经是引用了
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	//incoming request valiation 需要对request的类型验证
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}
	// if all is well , build the domain object & save the transaction
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.repo.SavaTransaction(t)
	if appError != nil {
		return nil, appError
	}

	response := transaction.ToDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo} //repo
}
