package model

type KeyValue[K any, V any] struct {
	Key   K `json:"key"`
	Value V `json:"value"`
}
