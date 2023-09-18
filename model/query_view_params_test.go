package model

import (
	"github.com/IosifSuzuki/go-couch_db/tool"
	"testing"
)

func TestQueryViewParams_BuildUrlValues(t *testing.T) {
	t.Run("queryViewParams string parameter", func(tt *testing.T) {
		params := QueryViewParams{
			Key: "filter",
		}
		queryViewParams, err := tool.BuildUrlValues(params)
		if err != nil {
			t.Errorf("build url failed: %v", err)
		} else if "key=filter" != queryViewParams.Encode() {
			t.Errorf("queryViewParams produce unexpected query text: %v", queryViewParams.Encode())
		}
	})

	t.Run("queryViewParams bool parameter", func(tt *testing.T) {
		params := QueryViewParams{
			Descending: true,
		}
		queryViewParams, err := tool.BuildUrlValues(params)
		if err != nil {
			t.Errorf("build url failed: %v", err)
		} else if "descending=true" != queryViewParams.Encode() {
			t.Errorf("queryViewParams produce unexpected query text: %v", queryViewParams.Encode())
		}
	})

	t.Run("queryViewParams int parameter", func(tt *testing.T) {
		params := QueryViewParams{
			Skip: 10,
		}
		queryViewParams, err := tool.BuildUrlValues(params)
		if err != nil {
			t.Errorf("build url failed: %v", err)
		} else if "skip=10" != queryViewParams.Encode() {
			t.Errorf("queryViewParams produce unexpected query text: %v", queryViewParams.Encode())
		}
	})
}
