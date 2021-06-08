package jxsaver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJXSaver_Validate(t *testing.T) {
	tcs := []struct {
		Name     string
		Format   string
		Data     string
		ExpError bool
	}{
		{
			Name:   "valid JSON data",
			Format: jsonFormat,
			Data:   `{"testKey1":"testValue1"}`,
		},
		{
			Name:     "invalid JSON data",
			Format:   jsonFormat,
			Data:     `{key: obviously invalid JSON}`,
			ExpError: true,
		},
		{
			Name:     "valid XML data",
			Format:   xmlFormat,
			Data:     `<thing><key1>value1</key1><key2>value2</key2></thing>`,
			ExpError: false,
		},
		{
			Name:     "invalid XML data",
			Format:   xmlFormat,
			Data:     `<thing><key1>value1</key1><key2>value2<key2></thing>`,
			ExpError: true,
		},
		{
			Name:     "invalid format",
			Format:   "invalid format",
			Data:     `some data of invalid format`,
			ExpError: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			err := NewApp().Validate(tc.Format, tc.Data)

			assert.Equal(t, tc.ExpError, err != nil)
		})
	}
}
