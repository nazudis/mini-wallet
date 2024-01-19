package transformer

import (
	"time"

	"github.com/nazudis/mini-wallet/engine/restapi/dto"
	"github.com/nazudis/mini-wallet/src/entity"
)

func TransformResponseWallet(wallet *entity.Wallet) (data dto.WalletResponseDto) {
	if wallet != nil {
		data = dto.WalletResponseDto{
			Id:      wallet.ID.String(),
			OwnedBy: wallet.OwnedBy.String(),
			Status:  wallet.Status,
			Balance: wallet.Balance.InexactFloat64(),
		}
		if wallet.Status == entity.WalletStatusEnabled {
			data.EnabledAt = wallet.EnabledAt.Format(time.RFC3339)
		}
		if wallet.Status == entity.WalletStatusDisabled {
			data.DisabledAt = wallet.DisabledAt.Format(time.RFC3339)
		}
	}

	return
}
