package model

type FindResponse struct {
	Docs []map[string]interface{} `json:"docs"`
}
