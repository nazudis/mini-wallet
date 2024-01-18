package controller

import (
	"net/http"

	"github.com/nazudis/mini-wallet/src/helper"
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

func NewWalletController(service service.WalletService) WalletController {
	return WalletControllerImpl{
		Service: service,
	}
}
