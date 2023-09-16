package logger

type LogLevel int

const (
	CriticalLevel LogLevel = iota + 1
	ErrorLevel
	WarningLevel
	NoticeLevel
	InfoLevel
	DebugLevel
)

func (l LogLevel) String() string {
	return []string{
		"CRI",
		"ERR",
		"WAR",
		"NOT",
		"INF",
		"DEB",
	}[l-1]
}

func (l LogLevel) Color() Color {
	var color Color
	switch l {
	case CriticalLevel:
		color = Magenta
	case ErrorLevel:
		color = Red
	case WarningLevel:
		color = Yellow
	case NoticeLevel:
		color = Green
	case InfoLevel:
		color = White
	case DebugLevel:
		color = Cyan
	}
	return color
}
