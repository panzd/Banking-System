package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Banking-System/dto"
	"github.com/Banking-System/errs"
	"github.com/gorilla/mux"

	"github.com/Banking-System/mocks/service"
	"github.com/golang/mock/gomock"
)

var router *mux.Router
var ch CustomerHandlers
var mockService *service.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	// 结束控制
	defer ctrl.Finish()

	mockService := service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{mockService}

	// 新建路由
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}
func Test_should_return_customers_with_status_code_200(t *testing.T) {

	teardown := setup(t)
	defer teardown()
	// Arrange

	dummyCustomers := []dto.CustomerResponse{
		{Id: "1001", Name: "Archie", City: "Shanghai", Zipcode: "210000", DateofBirth: "2022-01-01", Status: "1"},
		//status 活跃程度
	}

	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedError("some data error"))
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	//Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
