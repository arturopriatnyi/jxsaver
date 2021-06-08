package jxsaver

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"strings"
)

const (
	// jsonFormat indicates data in JSON format.
	jsonFormat = "json"
	// xmlFormat indicates data in XML format.
	xmlFormat = "xml"
)

// ErrInvalidFormat is returned when user provides format that is not "json" or "xml".
var ErrInvalidFormat = errors.New("invalid format provided")

// App is the main app interface with all business logic.
type App interface {
	Validate(format, data string) error
}

type app struct{}

// NewApp creates and return a new App instance.
func NewApp() App {
	return &app{}
}

// Validate validates data using provided format.
func (app *app) Validate(format, data string) error {
	r := strings.NewReader(data)

	var parsedData interface{}
	if format == jsonFormat {
		return json.NewDecoder(r).Decode(&parsedData)
	} else if format == xmlFormat {
		return xml.NewDecoder(r).Decode(&parsedData)
	}

	return ErrInvalidFormat
}
