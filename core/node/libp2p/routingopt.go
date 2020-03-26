package libp2p

import (
	"context"

	"github.com/ipfs/go-datastore"
	nilrouting "github.com/ipfs/go-ipfs-routing/none"
	host "github.com/libp2p/go-libp2p-core/host"
	routing "github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	record "github.com/libp2p/go-libp2p-record"
)

type RoutingOption func(context.Context, host.Host, datastore.Batching, record.Validator) (routing.Routing, error)

func constructDHTRouting(ctx context.Context, host host.Host, dstore datastore.Batching, validator record.Validator) (routing.Routing, error) {
	return dht.New(
		ctx, host,
		dht.Concurrency(10),
		dht.Mode(dht.ModeAuto),
		dht.Datastore(dstore),
		dht.Validator(validator),
	)
}

func constructClientDHTRouting(ctx context.Context, host host.Host, dstore datastore.Batching, validator record.Validator) (routing.Routing, error) {
	return dht.New(
		ctx, host,
		dht.Concurrency(10),
		dht.Mode(dht.ModeClient),
		dht.Datastore(dstore),
		dht.Validator(validator),
	)
}

var DHTOption RoutingOption = constructDHTRouting
var DHTClientOption RoutingOption = constructClientDHTRouting
var NilRouterOption RoutingOption = nilrouting.ConstructNilRouting
