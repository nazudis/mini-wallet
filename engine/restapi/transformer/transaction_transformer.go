package transformer

import (
	"time"

	"github.com/nazudis/mini-wallet/engine/restapi/dto"
	"github.com/nazudis/mini-wallet/src/entity"
	"github.com/nazudis/mini-wallet/src/helper"
)

func TransformResponseTransaction(transaction *entity.Transaction) (data dto.TransactionResponseDto) {
	if transaction != nil {
		trxDto := dto.Transaction{
			Id:           transaction.ID.String(),
			Status:       transaction.Status,
			TransactedAt: transaction.TransactedAt.Format(time.RFC3339),
			Type:         transaction.Type,
			Amount:       transaction.Amount.InexactFloat64(),
			ReferenceId:  transaction.ReferenceId.String(),
		}
		if transaction.Type == entity.TrxDeposit {
			trxDto.DepositedAt = helper.VarToPointer(transaction.TransactedAt.Format(time.RFC3339))
			trxDto.DepositedBy = helper.VarToPointer(transaction.OwnedBy.String())
			data.Deposit = &trxDto
		} else if transaction.Type == entity.TrxWithdrawal {
			trxDto.WithdrawnAt = helper.VarToPointer(transaction.TransactedAt.Format(time.RFC3339))
			trxDto.WithdrawnBy = helper.VarToPointer(transaction.OwnedBy.String())
			data.Withdrawal = &trxDto
		}
	}

	return
}

func TransformResponseTransactionList(transactions []entity.Transaction) (data dto.TransactionListResponseDto) {
	data = dto.TransactionListResponseDto{Transactions: []dto.Transaction{}}
	for _, trx := range transactions {
		trxDto := dto.Transaction{
			Id:           trx.ID.String(),
			Status:       trx.Status,
			TransactedAt: trx.TransactedAt.Format(time.RFC3339),
			Type:         trx.Type,
			Amount:       trx.Amount.InexactFloat64(),
			ReferenceId:  trx.ReferenceId.String(),
		}

		data.Transactions = append(data.Transactions, trxDto)
	}
	return
}
