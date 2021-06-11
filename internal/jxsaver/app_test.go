package jxsaver

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mockhasher "github.com/imarrche/jxsaver/pkg/hash/mock"
	mockstorage "github.com/imarrche/jxsaver/pkg/storage/mock"
)

func TestJXSaver_Init(t *testing.T) {
	tcs := []struct {
		Name         string
		Mock         func(c *gomock.Controller, sm *mockstorage.MockManager, hm *mockhasher.MockManager)
		ExpError     bool
		ExpHashes    []string
		ExpHashesLen int
	}{
		{
			Name: "init without errors, hashes file exists",
			Mock: func(c *gomock.Controller, sm *mockstorage.MockManager, hm *mockhasher.MockManager) {
				sm.EXPECT().FileExists(hashesFileName).Return(true)
				sm.EXPECT().ReadLinesFromFile(hashesFileName).Return([]string{"hash1", "hash2"}, nil)
			},
			ExpError:  false,
			ExpHashes: []string{"hash1", "hash2"},
		},
		{
			Name: "init without errors, hashes file doesn't exist",
			Mock: func(c *gomock.Controller, sm *mockstorage.MockManager, hm *mockhasher.MockManager) {
				sm.EXPECT().FileExists(hashesFileName).Return(false)
				sm.EXPECT().CreateFile(hashesFileName).Return(nil)
			},
			ExpError:  false,
			ExpHashes: []string{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			sm := mockstorage.NewMockManager(c)
			hm := mockhasher.NewMockManager(c)
			tc.Mock(c, sm, hm)

			app := NewApp(sm, hm)
			err := app.Init()
			hashes := make([]string, 0, len(app.Hashes))
			for k := range app.Hashes {
				hashes = append(hashes, k)
			}

			assert.Equal(t, tc.ExpError, err != nil)
			assert.Equal(t, tc.ExpHashes, hashes)
		})
	}
}

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
			Data:   `{"testKey1":"testValue1","testKey2":{"nestedKey1":["testValue2,testValue2"]}}`,
		},
		{
			Name:     "invalid JSON data",
			Format:   jsonFormat,
			Data:     `{key: obviously an invalid JSON}`,
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
			c := gomock.NewController(t)
			defer c.Finish()
			sm := mockstorage.NewMockManager(c)
			hm := mockhasher.NewMockManager(c)

			err := NewApp(sm, hm).Validate(tc.Format, tc.Data)

			assert.Equal(t, tc.ExpError, err != nil)
		})
	}
}

func TestJXSaver_Save(t *testing.T) {
	testHash := [16]byte{}
	testHashStr := "\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000"

	tcs := []struct {
		Name     string
		Mock     func(c *gomock.Controller, sm *mockstorage.MockManager, hm *mockhasher.MockManager, format string, data string)
		Format   string
		Data     string
		ExpError bool
	}{
		{
			Name: "data is saved",
			Mock: func(c *gomock.Controller, sm *mockstorage.MockManager, hm *mockhasher.MockManager, format string, data string) {
				hm.EXPECT().Hash([]byte(data)).Return(testHash)
				sm.EXPECT().WriteToFile(fmt.Sprintf("0.%s", format), data).Return(nil)
				sm.EXPECT().WriteToFile(hashesFileName, fmt.Sprintf("%s\n", testHashStr))
			},
			Format: jsonFormat,
			Data:   `{"testKey1":"testValue1"}`,
		},
		{
			Name: "duplicate data provided",
			Mock: func(c *gomock.Controller, sm *mockstorage.MockManager, hm *mockhasher.MockManager, format string, data string) {
				hm.EXPECT().Hash([]byte(data)).Return(testHash)
				sm.EXPECT().WriteToFile(fmt.Sprintf("0.%s", format), data).Return(nil)
				sm.EXPECT().WriteToFile(hashesFileName, fmt.Sprintf("%s\n", testHashStr))
			},
			Format: jsonFormat,
			Data:   `{"testKey1":"testValue1"}`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			sm := mockstorage.NewMockManager(c)
			hm := mockhasher.NewMockManager(c)
			tc.Mock(c, sm, hm, tc.Format, tc.Data)

			err := NewApp(sm, hm).Save(tc.Format, tc.Data)

			assert.Equal(t, tc.ExpError, err != nil)
		})
	}
}
