package domain

import (
	"strconv"

	"github.com/Banking-System/errs"
	"github.com/Banking-System/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {

	sqlInsert := "INSERT INTO accounts(customer_id, opening_date, account_type, amount, status)  values(?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while creating new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting last insert id for new account:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

func (d AccountRepositoryDb) SavaTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// starting the database transaction block
	tx, err := d.client.Begin() //Begin starts a transaction.
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	//inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO transaction(account_id, amount, transaction_type,transaction_date)
													values(?,?,?,?)`,
		t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// in case of error Rollback, and changes from both the tables will be reverted
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId) // - 去
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}
	//commit the transaction when all  is good
	if err != nil {
		tx.Rollback() //回滚
		logger.Error("Error while saving transaction:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	//getting the last transaction ID from the transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	//getting the latest account information from  the accounts table
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10) //十进制的

	//updating the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date,account_type, amount from accounts where account_id=?"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}
