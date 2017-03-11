// Package syncosc exists to define constants used in the oscsync protocol.
// See http://github.com/scgolang/oscsync/README.md
package syncosc

// OSC addresses.
const (
	AddressPulse       = "/sync/pulse"
	AddressSlaveAdd    = "/sync/slave/add"
	AddressSlaveList   = "/sync/slave/list"
	AddressSlaveRemove = "/sync/slave/remove"
	AddressTempo       = "/sync/tempo"
)

// PulsesPerBar is the number of pulses in a bar (measure).
const PulsesPerBar = 96
