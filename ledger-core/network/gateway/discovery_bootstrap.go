package gateway

import (
	"context"
	"time"

	"github.com/insolar/assured-ledger/ledger-core/appctl/beat"
	"github.com/insolar/assured-ledger/ledger-core/insolar/nodeinfo"
	"github.com/insolar/assured-ledger/ledger-core/instrumentation/inslogger"
	"github.com/insolar/assured-ledger/ledger-core/pulse"
)

func newDiscoveryBootstrap(b *Base) *DiscoveryBootstrap {
	return &DiscoveryBootstrap{b}
}

// DiscoveryBootstrap void network state
type DiscoveryBootstrap struct {
	*Base
}

func (g *DiscoveryBootstrap) Run(ctx context.Context, p pulse.Data) {
	logger := inslogger.FromContext(ctx)
	cert := g.CertificateManager.GetCertificate()

	permit, err := g.BootstrapRequester.Authorize(ctx, cert)
	if err != nil {
		logger.Warn("Failed to authorize: ", err.Error())
		g.Gatewayer.SwitchState(ctx, nodeinfo.NoNetworkState, p)
		return
	}

	resp, err := g.BootstrapRequester.Bootstrap(ctx, permit, *g.originCandidate)
	if err != nil {
		logger.Warn("Failed to bootstrap: ", err.Error())
		g.Gatewayer.SwitchState(ctx, nodeinfo.NoNetworkState, p)
		return
	}

	logger.Infof("Bootstrapping to node %s", permit.Payload.ReconnectTo)

	// Reset backoff if not insolar.NoNetworkState.
	g.backoff.Reset()

	// no needed
	responsePulse := beat.Beat{}

	g.bootstrapETA = time.Second * time.Duration(resp.ETASeconds)
	g.bootstrapTimer = time.NewTimer(g.bootstrapETA)
	g.Gatewayer.SwitchState(ctx, nodeinfo.WaitConsensus, responsePulse.Data)
}

func (g *DiscoveryBootstrap) GetState() nodeinfo.NetworkState {
	return nodeinfo.DiscoveryBootstrap
}