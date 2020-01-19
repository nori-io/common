package logger

import "time"

type Hook interface {
	Levels() []Level
	Fire(e Entry, field ...Field) error
}

type Field struct {
	Key   string
	Value string
}

type Entry struct {
	Formatter Formatter
	Level     Level
	Time      time.Time
	Message   string
}

type Formatter interface {
	Format(e Entry, field ...Field) ([]byte, error)
}

type Logger interface {
	FieldLogger
	AddHook(hook Hook)
	With(fields ...Field) Logger
}

type FieldLogger interface {
	// Critical push to log entry with critical level
	Critical(format string, opts ...interface{})
	// Debug push to log entry with debug level
	Debug(format string, opts ...interface{})
	// Fatal logs a message with fatal level and exit with status set to 1
	Fatal(format string, opts ...interface{})
	// Error push to log entry with error level
	Error(format string, opts ...interface{})
	// Info push to log entry with info level
	Info(format string, opts ...interface{})
	// Log push to log with specified level
	Log(level Level, format string, opts ...interface{})
	// Notice push to log entry with notice level
	Notice(format string, opts ...interface{})
	// Panic logs a message with panic level and then throws the panic
	Panic(format string, opts ...interface{})
	// Warning push to log entry with warning level
	Warning(format string, opts ...interface{})
}
