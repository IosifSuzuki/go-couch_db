package model

import "fmt"

type ErrorResponse struct {
	Description string `json:"error"`
	Reason      string `json:"reason"`
	Code        int
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("status code: %d; error: %s; reason: %s", e.Code, e.Description, e.Reason)
}
