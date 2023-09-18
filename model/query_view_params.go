package model

type QueryViewParams struct {
	Key        string `query:"key,omitempty"`
	Descending bool   `query:"descending,omitempty"`
	Skip       int    `query:"skip,omitempty"`
}
