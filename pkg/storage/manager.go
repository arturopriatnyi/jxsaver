// Package storage provides utilities to work with file system.
package storage

import (
	"bufio"
	"os"
)

//go:generate mockgen -source=manager.go -destination=mock/mock.go

// Manager is interface for performing operations with file system.
type Manager interface {
	FileExists(fileName string) bool
	CreateFile(name string) error
	WriteToFile(fileName, content string) error
	ReadLinesFromFile(fileName string) ([]string, error)
}

// manager is Manager implementation.
type manager struct{}

// NewManager creates and returns a new Manager instance.
func NewManager() Manager {
	return &manager{}
}

// FileExists checks if file already exists.
func (m *manager) FileExists(fileName string) bool {
	_, err := os.Stat(fileName)

	return err == nil
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

// ReadLinesFromFile reads all lines of a file to slice of strings.
func (m *manager) ReadLinesFromFile(fileName string) (lines []string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return lines, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
