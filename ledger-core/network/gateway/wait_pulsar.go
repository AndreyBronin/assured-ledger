// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package gateway

import (
	"context"

	"github.com/insolar/assured-ledger/ledger-core/instrumentation/inslogger"
	"github.com/insolar/assured-ledger/ledger-core/network"
	"github.com/insolar/assured-ledger/ledger-core/network/consensus/gcpv2/api/census"
	"github.com/insolar/assured-ledger/ledger-core/network/consensus/gcpv2/api/proofs"
	"github.com/insolar/assured-ledger/ledger-core/network/rules"
	"github.com/insolar/assured-ledger/ledger-core/pulse"
	"github.com/insolar/assured-ledger/ledger-core/vanilla/throw"
)

func newWaitPulsar(b *Base) *WaitPulsar {
	return &WaitPulsar{b, make(chan pulse.Data, 1)}
}

type WaitPulsar struct {
	*Base
	pulseArrived chan pulse.Data
}

func (g *WaitPulsar) Run(ctx context.Context, pulse pulse.Data) {
	g.switchOnRealPulse(pulse)

	select {
	case <-g.bootstrapTimer.C:
		g.FailState(ctx, bootstrapTimeoutMessage)
	case newPulse := <-g.pulseArrived:
		g.Gatewayer.SwitchState(ctx, network.CompleteNetworkState, newPulse)
	}
}

func (g *WaitPulsar) UpdateState(ctx context.Context, pulseNumber pulse.Number, isTimePulse bool, nodes census.OnlinePopulation, csh proofs.CloudStateHash) {
	if _, err := rules.CheckMajorityRule(g.CertificateManager.GetCertificate(), nodes); err != nil {
		g.FailState(ctx, err.Error())
	}

	if err := rules.CheckMinRole(g.CertificateManager.GetCertificate(), nodes); err != nil {
		g.FailState(ctx, err.Error())
	}

	g.Base.UpdateState(ctx, pulseNumber, isTimePulse, nodes, csh)
}

func (g *WaitPulsar) GetState() network.State {
	return network.WaitPulsar
}

func (g *WaitPulsar) OnConsensusFinished(ctx context.Context, report network.Report) {
	g.switchOnRealPulse(EnsureGetPulse(ctx, report))
}

func (g *WaitPulsar) switchOnRealPulse(pulseObject pulse.Data) {
	if pulseObject.PulseNumber.IsTimePulse() && pulseObject.PulseEpoch.IsTimeEpoch() {
		g.pulseArrived <- pulseObject
		close(g.pulseArrived)
	}
}

func (g *WaitPulsar) OnPulseFromConsensus(ctx context.Context, pulse network.NetworkedPulse) {
	g.Base.OnPulseFromConsensus(ctx, pulse)

	if !pulse.PulseEpoch.IsTimeEpoch() {
		panic(throw.IllegalState())
	}

	err := g.PulseAppender.Append(ctx, pulse)
	if err != nil {
		inslogger.FromContext(ctx).Panic("failed to append pulse: ", err.Error())
	}

	err = g.PulseManager.CommitFirstPulseChange(pulse)
	if err != nil {
		inslogger.FromContext(ctx).Panicf("failed to set start pulse: %d, %s", pulse.PulseNumber, err.Error())
	}
}
