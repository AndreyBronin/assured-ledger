// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package testwallet

import (
	"github.com/insolar/assured-ledger/ledger-core/v2/logicrunner/builtin/foundation"
)

// Wallet - basic wallet contract.
type Wallet struct {
	foundation.BaseContract
	balance uint32
}

const initialBalance = 1000000000

// New creates new wallet.
func New() (*Wallet, error) {
	return &Wallet{balance: initialBalance}, nil
}

func (w *Wallet) Balance() (uint32, error) {
	return w.balance, nil
}

func (w *Wallet) Accept(amount uint32) error {
	w.balance += amount
	return nil
}
