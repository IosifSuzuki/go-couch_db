package model

type ViewResponse struct {
	TotalRows *int             `json:"total_rows"`
	Offset    *int             `json:"offset"`
	Rows      []map[string]any `json:"rows"`
}
