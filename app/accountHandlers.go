package app

import (
	"banking/dto"
	"banking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_id := vars["customer_id"]
	var request dto.NewAccountRequest

	errs := json.NewDecoder(r.Body).Decode(&request)
	if errs != nil {
		writeResponse(w, http.StatusBadRequest, errs.Error())
	} else {
		request.CustomerId = customer_id
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

// /customers/{customer_id}/accounts/{account_id}
func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	// decode incoming request
	var request dto.TransactionRequest
	errs := json.NewDecoder(r.Body).Decode(&request)
	if errs != nil {
		writeResponse(w, http.StatusBadRequest, errs.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId

		transaction, appError := h.service.MakeTransaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusOK, transaction)
		}
	}
}
