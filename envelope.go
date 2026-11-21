// Package contract is the shared vocabulary and message set between the IoT
// gateway and the backend. It holds types and constants only: no transport and
// no business logic.
package contract

import (
	"encoding/json"
	"time"
)

// Kind names the body carried by an Envelope.
type Kind string

const (
	// KindInventory is a full device inventory snapshot (upward).
	KindInventory Kind = "inventory"
	// KindConfig is a whole-kind configuration document (downward).
	KindConfig Kind = "config"
	// KindConfigReport is a full report of the configuration a gateway holds (upward).
	KindConfigReport Kind = "config_report"
)

// Envelope wraps every message, upward and downward. Downward, the backend leaves
// BootID and Seq empty: ordering is decided from the upward (BootID, Seq) only.
type Envelope struct {
	GatewayID string          `json:"gateway_id"`
	BootID    string          `json:"boot_id"`
	Seq       uint64          `json:"seq"`
	SentAt    time.Time       `json:"sent_at"` // informational only, never used for decisions
	Kind      Kind            `json:"kind"`
	Body      json.RawMessage `json:"body"`
}

// NewerThan reports whether e is newer than the last processed (bootID, seq).
// A different boot is always newer; within one boot only a higher seq is.
func (e Envelope) NewerThan(bootID string, seq uint64) bool {
	if e.BootID != bootID {
		return true
	}
	return e.Seq > seq
}

// Wrap marshals body into an Envelope stamped with the current time.
func Wrap(gatewayID, bootID string, seq uint64, kind Kind, body any) (Envelope, error) {
	raw, err := json.Marshal(body)
	if err != nil {
		return Envelope{}, err
	}
	return Envelope{GatewayID: gatewayID, BootID: bootID, Seq: seq, SentAt: time.Now().UTC(), Kind: kind, Body: raw}, nil
}
