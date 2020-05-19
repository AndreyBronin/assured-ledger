// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package merkle

import (
	"github.com/insolar/assured-ledger/ledger-core/v2/cryptography"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar/node"
)

type BaseProof struct {
	Signature cryptography.Signature
}

func (bp *BaseProof) signature() cryptography.Signature {
	return bp.Signature
}

type PulseProof struct {
	BaseProof

	StateHash []byte
}

func (np *PulseProof) hash(pulseHash []byte, helper *merkleHelper) []byte {
	return helper.nodeInfoHash(pulseHash, np.StateHash)
}

type GlobuleProof struct {
	BaseProof

	PrevCloudHash []byte
	GlobuleID     node.GlobuleID
	NodeCount     uint32
	NodeRoot      []byte
}

func (gp *GlobuleProof) hash(globuleHash []byte, _ *merkleHelper) []byte {
	return globuleHash
}

type CloudProof struct {
	BaseProof
}

func (cp *CloudProof) hash(cloudHash []byte, _ *merkleHelper) []byte {
	return cloudHash
}
