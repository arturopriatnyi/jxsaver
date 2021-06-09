// Package hash provides utilities to work with hashing.
package hash

import "crypto/md5"

//go:generate mockgen -source=manager.go -destination=mock/mock.go

// Manager is interface for performing hashing operations.
type Manager interface {
	Hash([]byte) [16]byte
}

// manager is Manager implementation.
type manager struct{}

// NewManager creates and returns a new Manager instance.
func NewManager() Manager {
	return &manager{}
}

// Hash returns hash sum of a data.
func (m *manager) Hash(data []byte) [16]byte {
	return md5.Sum(data)
}
