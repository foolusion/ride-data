// Package ridefile implements a simple ride file interface for reading
// exported ride data.
package ridefile

type RideFile interface {
	Read() []byte, error
}
