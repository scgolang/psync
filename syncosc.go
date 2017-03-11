// Package syncosc exists to define constants used in the oscsync protocol.
// See http://github.com/scgolang/oscsync/README.md
package syncosc

import (
	"time"
)

// OSC addresses.
const (
	AddressPulse       = "/sync/pulse"
	AddressSlaveAdd    = "/sync/slave/add"
	AddressSlaveList   = "/sync/slave/list"
	AddressSlaveRemove = "/sync/slave/remove"
	AddressTempo       = "/sync/tempo"
)

// MasterPort is the listening port for the oscsync master.
const MasterPort = 5776

// PulsesPerBar is the number of pulses in a bar (measure).
const PulsesPerBar = 96

// GetPulseDuration converts the tempo in bpm to a time.Duration
// callers are responsible for making concurrent access safe.
func GetPulseDuration(tempo float32) time.Duration {
	if tempo == 0 {
		return time.Duration(0)
	}
	return time.Duration(float32(int64(24e10)/PulsesPerBar) / tempo)
}
