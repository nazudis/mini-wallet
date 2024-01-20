package controller

import (
	"net/http"

	"github.com/nazudis/mini-wallet/engine/restapi/transformer"
	"github.com/nazudis/mini-wallet/src/helper"
	"github.com/nazudis/mini-wallet/src/middleware"
	"github.com/nazudis/mini-wallet/src/service"
	"github.com/shopspring/decimal"
)

type WalletControllerImpl struct {
	Service service.WalletService
}

// InitAccount implements WalletController.
func (c WalletControllerImpl) InitAccount(w http.ResponseWriter, r *http.Request) {
	res := helper.PlugResponse(w)

	customerXid := r.FormValue("customer_xid")
	if customerXid == "" {
		data := helper.MapJSON{"error": "customer_xid is required"}
		res.SetHttpStatusCode(http.StatusBadRequest).ReplyFail(data)
		return
	}

	token, err := c.Service.InitAccount(customerXid)
	if err != nil {
		data := helper.MapJSON{"error": helper.MutatedValue(err.Error())}
		res.SetHttpStatusCode(http.StatusBadRequest).ReplyFail(data)
		return
	}

	data := helper.MapJSON{"token": helper.MutatedValue(token)}
	res.ReplySuccess(data)
}

// EnableWalletAccount implements WalletController.
func (c WalletControllerImpl) EnableWalletAccount(w http.ResponseWriter, r *http.Request) {
	res := helper.PlugResponse(w)

	customerXid := r.Context().Value(middleware.CustomerXidCtxKey).(string)

	wallet, err := c.Service.EnableWalletAccount(customerXid)
	if err != nil {
		data := helper.MapJSON{"error": helper.MutatedValue(err.Error())}
		res.SetHttpStatusCode(http.StatusBadRequest).ReplyFail(data)
		return
	}

	data := transformer.TransformResponseWallet(wallet)
	res.ReplySuccess(data)
}

// DisableWalletAccount implements WalletController.
func (c WalletControllerImpl) DisableWalletAccount(w http.ResponseWriter, r *http.Request) {
	res := helper.PlugResponse(w)

	customerXid := r.Context().Value(middleware.CustomerXidCtxKey).(string)
	isDisabled := r.FormValue("is_disabled")

	wallet, err := c.Service.DisableWalletAccount(customerXid, isDisabled == "true")
	if err != nil {
		data := helper.MapJSON{"error": helper.MutatedValue(err.Error())}
		res.SetHttpStatusCode(http.StatusBadRequest).ReplyFail(data)
		return
	}

	data := transformer.TransformResponseWallet(wallet)
	res.ReplySuccess(data)
}

// GetWallet implements WalletController.
func (c WalletControllerImpl) GetWallet(w http.ResponseWriter, r *http.Request) {
	res := helper.PlugResponse(w)

	customerXid := r.Context().Value(middleware.CustomerXidCtxKey).(string)

	wallet, err := c.Service.GetEnabledWallet(customerXid)
	if err != nil {
		data := helper.MapJSON{"error": helper.MutatedValue(err.Error())}
		res.SetHttpStatusCode(http.StatusBadRequest).ReplyFail(data)
		return
	}

	data := transformer.TransformResponseWallet(wallet)
	res.ReplySuccess(data)
}

// Deposit implements WalletController.
func (c WalletControllerImpl) Deposit(w http.ResponseWriter, r *http.Request) {
	res := helper.PlugResponse(w)

	customerXid := r.Context().Value(middleware.CustomerXidCtxKey).(string)

	amountString := r.FormValue("amount")
	amount, err := decimal.NewFromString(amountString)
	if err != nil {
		data := helper.MapJSON{"error": helper.MutatedValue(err.Error())}
		res.SetHttpStatusCode(http.StatusBadRequest).ReplyFail(data)
		return
	}

	referenceId := r.FormValue("reference_id")
	if referenceId == "" {
		data := helper.MapJSON{"error": "reference_id is required"}
		res.SetHttpStatusCode(http.StatusBadRequest).ReplyFail(data)
		return
	}

	trx, err := c.Service.Deposit(service.TransactionParams{
		CustomerXid: customerXid,
		Amount:      amount,
		ReferenceId: referenceId,
	})
	if err != nil {
		data := helper.MapJSON{"error": helper.MutatedValue(err.Error())}
		res.SetHttpStatusCode(http.StatusBadRequest).ReplyFail(data)
		return
	}

	data := transformer.TransformResponseTransaction(trx)
	res.ReplySuccess(data)
}

// Withdraw implements WalletController.
func (c WalletControllerImpl) Withdraw(w http.ResponseWriter, r *http.Request) {
	res := helper.PlugResponse(w)

	customerXid := r.Context().Value(middleware.CustomerXidCtxKey).(string)

	amountString := r.FormValue("amount")
	amount, err := decimal.NewFromString(amountString)
	if err != nil {
		data := helper.MapJSON{"error": helper.MutatedValue(err.Error())}
		res.SetHttpStatusCode(http.StatusBadRequest).ReplyFail(data)
		return
	}

	referenceId := r.FormValue("reference_id")
	if referenceId == "" {
		data := helper.MapJSON{"error": "reference_id is required"}
		res.SetHttpStatusCode(http.StatusBadRequest).ReplyFail(data)
		return
	}

	trx, err := c.Service.Withdraw(service.TransactionParams{
		CustomerXid: customerXid,
		Amount:      amount,
		ReferenceId: referenceId,
	})
	if err != nil {
		data := helper.MapJSON{"error": helper.MutatedValue(err.Error())}
		res.SetHttpStatusCode(http.StatusBadRequest).ReplyFail(data)
		return
	}

	data := transformer.TransformResponseTransaction(trx)
	res.ReplySuccess(data)
}

// GetWalletTransactions implements WalletController.
func (c WalletControllerImpl) GetWalletTransactions(w http.ResponseWriter, r *http.Request) {
	res := helper.PlugResponse(w)

	customerXid := r.Context().Value(middleware.CustomerXidCtxKey).(string)

	trxs, err := c.Service.GetTransactions(customerXid)
	if err != nil {
		data := helper.MapJSON{"error": helper.MutatedValue(err.Error())}
		res.SetHttpStatusCode(http.StatusBadRequest).ReplyFail(data)
		return
	}

	data := transformer.TransformResponseTransactionList(trxs)
	res.ReplySuccess(data)
}

func NewWalletController(service service.WalletService) WalletController {
	return WalletControllerImpl{
		Service: service,
	}
}
