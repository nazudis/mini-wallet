package controller

import (
	"net/http"

	"github.com/nazudis/mini-wallet/engine/restapi/transformer"
	"github.com/nazudis/mini-wallet/src/helper"
	"github.com/nazudis/mini-wallet/src/middleware"
	"github.com/nazudis/mini-wallet/src/service"
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

func NewWalletController(service service.WalletService) WalletController {
	return WalletControllerImpl{
		Service: service,
	}
}
