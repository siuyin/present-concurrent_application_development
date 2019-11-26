// Package order receives, validates and sends orders
package order

type Event int

//go:generate stringer -type=Event

// order events -- must be append only. Do not change sequence.
const (
	Created Event = iota + 1
	Invalid
	Completed
	Fulfilled
)
