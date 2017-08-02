package log

import "github.com/fatih/color"

// LogSimple - prints white message
func LogSimple(format string, args ...interface{}) {
	color.White(format, args...)
}

// Deprecated: use LogWarning
// LogSkip - print yellow message
func LogSkip(format string, args ...interface{}) {
	color.Yellow(format, args...)
}

// LogWarning - prints yellow message
func LogWarning(format string, args ...interface{}) {
	LogSkip(format, args)
}

// LogError - prints red error message
func LogError(format string, args ...interface{}) {
	color.Red(format, args...)
}

// LogSuccess - prints green success message
func LogSuccess(format string, args ...interface{}) {
	color.Green(format, args...)
}
