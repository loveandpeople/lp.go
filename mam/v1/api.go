package mam

import (
	"github.com/loveandpeople/lp.go/api"
	"github.com/loveandpeople/lp.go/bundle"
	"github.com/loveandpeople/lp.go/transaction"
	"github.com/loveandpeople/lp.go/trinary"
)

// API defines an interface with a subset of methods of `api.API`.
type API interface {
	PrepareTransfers(seed trinary.Trytes, transfers bundle.Transfers, opts api.PrepareTransfersOptions) ([]trinary.Trytes, error)
	SendTrytes(trytes []trinary.Trytes, depth uint64, mwm uint64, reference ...trinary.Hash) (bundle.Bundle, error)
	FindTransactionObjects(query api.FindTransactionsQuery) (transaction.Transactions, error)
}
