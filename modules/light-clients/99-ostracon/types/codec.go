package types

import (
	codectypes "github.com/line/lbm-sdk/codec/types"

	"github.com/line/ibc-go/v3/modules/core/exported"
)

// RegisterInterfaces registers the ostracon concrete client-related
// implementations and interfaces.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*exported.ClientState)(nil),
		&ClientState{},
	)
	registry.RegisterImplementations(
		(*exported.ConsensusState)(nil),
		&ConsensusState{},
	)
	registry.RegisterImplementations(
		(*exported.Header)(nil),
		&Header{},
	)
	registry.RegisterImplementations(
		(*exported.Misbehaviour)(nil),
		&Misbehaviour{},
	)
}