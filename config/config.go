package config

var (
	logger *Logger
)

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
