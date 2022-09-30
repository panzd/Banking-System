package app

import (
	"encoding/json"
	"net/http"

	"github.com/Banking-System/dto"
	"github.com/Banking-System/service"
	"github.com/gorilla/mux"
)

type AccountHander struct {
	service service.AccountService
}

func (h AccountHander) NewAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r) // context 从请求中提取关键词，之前遇到过的

	customerId := vars["customer_id"]

	var request dto.NewAccountRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {

		request.CustomerId = customerId

		account, appError := h.service.NewAccount(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h AccountHander) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	//get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	//decode incoming request
	var request dto.TransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		// build the request object
		request.AccountId = accountId
		request.CustomerId = customerId

		//make transaction
		account, appError := h.service.MakeTransaction(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}
