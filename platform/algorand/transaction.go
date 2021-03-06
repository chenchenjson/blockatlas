package algorand

import (
	"strconv"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.TxPage, error) {
	txs, err := p.client.GetTxsOfAddress(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTxs(txs), nil
}

func NormalizeTxs(txs []Transaction) []types.Tx {
	result := make([]types.Tx, 0)

	for _, tx := range txs {
		if normalized, ok := Normalize(tx); ok {
			result = append(result, normalized)
		}
	}

	return result
}

func Normalize(tx Transaction) (result types.Tx, ok bool) {

	if tx.Type != TransactionTypePay {
		return result, false
	}

	return types.Tx{
		ID:     tx.Hash,
		Coin:   coin.ALGO,
		From:   tx.From,
		To:     tx.Payment.Receiver,
		Fee:    types.Amount(strconv.Itoa(int(tx.Fee))),
		Date:   int64(tx.Timestamp),
		Block:  tx.Round,
		Status: types.StatusCompleted,
		Type:   types.TxTransfer,
		Meta: types.Transfer{
			Value:    types.Amount(strconv.Itoa(int(tx.Payment.Amount))),
			Symbol:   coin.Coins[coin.ALGO].Symbol,
			Decimals: coin.Coins[coin.ALGO].Decimals,
		},
	}, true
}
