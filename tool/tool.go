package tool

func String(value string) *string {
	return &value
}

func IsSuccessStatusCode(code int) bool {
	return code >= 200 && code < 300
}
