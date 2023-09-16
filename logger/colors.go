package logger

import "fmt"

type Color int

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func (c Color) colorString() string {
	return fmt.Sprintf("\033[%dm", c)
}
