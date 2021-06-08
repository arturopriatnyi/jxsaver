// Package storage provides utilities to work with file system.
package storage

import "os"

//go:generate mockgen -source=manager.go -destination=mock/mock.go

// Manager is interface for performing operations with file system.
type Manager interface {
	CreateFile(name string) error
	WriteToFile(fileName, content string) error
}

// manager is Manager implementation.
type manager struct{}

// NewManager creates and returns a new Manager instance.
func NewManager() Manager {
	return &manager{}
}

// CreateFile creates a new file with provided name in current directory.
func (m *manager) CreateFile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

// WriteToFile appends provided content to the file with provided name, if file does not exist
// also creates it.
func (m *manager) WriteToFile(fileName, content string) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}
