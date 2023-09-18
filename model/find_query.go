package model

type Find[T any] struct {
	Selector T
	Fields   []string
	Sort     []string
	Limit    string
	Skip     string
}
