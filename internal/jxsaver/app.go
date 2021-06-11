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
	// ErrInvalidData is return when user provides invalid data.
	ErrInvalidData = errors.New("invalid data provided")
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

	// Hydrating hash store from file.
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
func (app *App) Validate(format, data string) (err error) {
	r := strings.NewReader(data)

	var parsedData interface{}
	if format == jsonFormat {
		err = json.NewDecoder(r).Decode(&parsedData)
	} else if format == xmlFormat {
		err = xml.NewDecoder(r).Decode(&parsedData)
	} else {
		return ErrInvalidFormat
	}

	if err != nil {
		return ErrInvalidData
	}
	return nil
}

// Save saves unique data to the file system.
func (app *App) Save(format, data string) error {
	// Calculating hash before validation to not validate duplicate data.
	hash := app.hasher.Hash([]byte(data))
	hashStr := string(hash[:])
	if app.Hashes[hashStr] {
		return ErrDuplicateData
	}

	if err := app.Validate(format, data); err != nil {
		return err
	}

	fileName := fmt.Sprintf("%d.%s", len(app.Hashes), format)
	if err := app.store.WriteToFile(fileName, data); err != nil {
		return err
	}
	app.Hashes[hashStr] = true

	// Saving hash to file, each hash on a new line.
	return app.store.WriteToFile(hashesFileName, fmt.Sprintf("%s\n", hashStr))
}
