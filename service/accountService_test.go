package service

import (
	"testing"

	realdomain "github.com/Banking-System/domain"
	"github.com/Banking-System/dto"
	"github.com/Banking-System/errs"
	"github.com/Banking-System/mocks/domain"
	"github.com/golang/mock/gomock"
)

func Test_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T) {
	request := dto.NewAccountRequest{
		CustomerId:  "2022",
		AccountType: "Saving",
		Amount:      200,
	}
	service := NewAccountService(nil)
	//Act
	_, appError := service.NewAccount(request)
	//Assert
	if appError == nil {
		t.Error("failed while testing the new account validation")
	}
}

var mockRepo *domain.MockAccountRepository
var service AccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}

}
func Test_should_return_an_error_from_the_server_side_if_the_new_account_cannnot_be_created(t *testing.T) {

	//Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "2022",
		AccountType: "Saving",
		Amount:      7000,
	}
	account := realdomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: dbTSLayout,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	mockRepo.EXPECT().Save(account).Return(nil, errs.NewUnexpectedError("unexpected"))
	// Act
	_, appError := service.NewAccount(req)

	if appError == nil {
		t.Error("Test failed")
	}
}
