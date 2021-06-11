package jxsaver

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"

	"github.com/imarrche/jxsaver/pkg/hash"
	"github.com/imarrche/jxsaver/pkg/storage"
)

const (
	// jsonFormat indicates data in JSON format.
	jsonFormat = "json"
	// xmlFormat indicates data in XML format.
	xmlFormat = "xml"
	// hashesFileName is name of file with file hashes.
	hashesFileName = "hashes.dat"
)

var (
	// ErrInvalidFormat is returned when user provides format that is not "json" or "xml".
	ErrInvalidFormat = errors.New("invalid format provided")
	// ErrDuplicateData is returned when user provides data that is already saved.
	ErrDuplicateData = errors.New("duplicate data provided")
)

type App struct {
	store  storage.Manager
	hasher hash.Manager
	Hashes map[string]bool
}

// NewApp creates and return a new App instance.
func NewApp(store storage.Manager, hasher hash.Manager) *App {
	app := &App{
		store:  store,
		hasher: hasher,
		Hashes: map[string]bool{},
	}

	return app
}

// Init performs all initializations steps to make app work.
func (app *App) Init() error {
	if !app.store.FileExists(hashesFileName) {
		return app.store.CreateFile(hashesFileName)
	}

	hashes, err := app.store.ReadLinesFromFile(hashesFileName)
	if err != nil {
		return err
	}

	for _, h := range hashes {
		app.Hashes[h] = true
	}

	return nil
}

// Validate validates data using provided format.
func (app *App) Validate(format, data string) error {
	r := strings.NewReader(data)

	var parsedData interface{}
	if format == jsonFormat {
		return json.NewDecoder(r).Decode(&parsedData)
	} else if format == xmlFormat {
		return xml.NewDecoder(r).Decode(&parsedData)
	}

	return ErrInvalidFormat
}

// Save saves unique data to the file system.
func (app *App) Save(format, data string) error {
	hash := app.hasher.Hash([]byte(data))
	hashStr := string(hash[:])
	if app.Hashes[hashStr] {
		return ErrDuplicateData
	}

	err := app.Validate(format, data)
	if err != nil {
		return err
	}
	if err = app.store.WriteToFile(fmt.Sprintf("%d.%s", len(app.Hashes), format), data); err != nil {
		return err
	}
	app.Hashes[hashStr] = true

	return app.store.WriteToFile(hashesFileName, fmt.Sprintf("%s\n", hashStr))
}
