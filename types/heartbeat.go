package types

import (
	"fmt"
	"io"

	"github.com/tendermint/go-crypto"
	"github.com/tendermint/go-wire"
	"github.com/tendermint/go-wire/data"
	cmn "github.com/tendermint/tmlibs/common"
)

// Heartbeat is a simple vote-like structure so validators can alert others that
// they are alive and waiting for transactions.
type Heartbeat struct {
	ValidatorAddress data.Bytes       `json:"validator_address"`
	ValidatorIndex   int              `json:"validator_index"`
	Height           uint64           `json:"height"`
	Round            int              `json:"round"`
	Sequence         int              `json:"sequence"`
	Signature        crypto.Signature `json:"signature"`
}

// WriteSignBytes writes the Heartbeat for signing.
func (heartbeat *Heartbeat) WriteSignBytes(chainID string, w io.Writer, n *int, err *error) {
	wire.WriteJSON(CanonicalJSONOnceHeartbeat{
		chainID,
		CanonicalHeartbeat(heartbeat),
	}, w, n, err)
}

// Copy makes a copy of the Heartbeat.
func (heartbeat *Heartbeat) Copy() *Heartbeat {
	heartbeatCopy := *heartbeat
	return &heartbeatCopy
}

// String returns a string representation of the Heartbeat.
func (heartbeat *Heartbeat) String() string {
	if heartbeat == nil {
		return "nil-heartbeat"
	}

	return fmt.Sprintf("Heartbeat{%v:%X %v/%02d (%v) %v}",
		heartbeat.ValidatorIndex, cmn.Fingerprint(heartbeat.ValidatorAddress),
		heartbeat.Height, heartbeat.Round, heartbeat.Sequence, heartbeat.Signature)
}
