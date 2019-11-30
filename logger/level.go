package logger

type Level uint8

const (
	// Debug: debug-level messages
	LevelDebug Level = 7
	// Informational: informational messages
	LevelInfo Level = 6
	// Warning: warning conditions. Non-critical entries that deserve eyes.
	LevelWarning Level = 5
	// Error: error conditions. Used for errors that should definitely be noted.
	LevelError Level = 4
	// Critical: critical conditions.
	LevelCritical Level = 3
	// Alert: action must be taken immediately.
	LevelAlert Level = 2
	// Panic: system panics. Logs and then calls panic with passed message.
	LevelPanic Level = 1
	// Fatal: system is unusable. Logs and then calls `os.Exit(1)`.
	LevelFatal Level = 0
)

const (
	// Debug: debug-level messages
	LevelDebugStr = "debug"
	// Informational: informational messages
	LevelInfoStr = "info"
	// Warning: warning conditions
	LevelWarningStr = "warning"
	// Error: error conditions
	LevelErrorStr = "error"
	// Critical: critical conditions
	LevelCriticalStr = "critical"
	// Alert: action must be taken immediately
	LevelAlertStr = "alert"
	// Panic: logger calls panic with passed message
	LevelPanicStr = "panic"
	// Fatal: system is unusable
	LevelFatalStr = "fatal"
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return LevelDebugStr
	case LevelInfo:
		return LevelInfoStr
	case LevelWarning:
		return LevelWarningStr
	case LevelError:
		return LevelErrorStr
	case LevelAlert:
		return LevelAlertStr
	case LevelCritical:
		return LevelCriticalStr
	case LevelFatal:
		return LevelAlertStr
	case LevelPanic:
		return LevelPanicStr
	}
	return "unknown"
}

func (l *Level) FromString(level string) Level {
	switch level {
	default:
		*l = LevelDebug
	case LevelDebugStr:
		*l = LevelDebug
	case LevelInfoStr:
		*l = LevelInfo
	case LevelWarningStr:
		*l = LevelWarning
	case LevelErrorStr:
		*l = LevelError
	case LevelAlertStr:
		*l = LevelAlert
	case LevelCriticalStr:
		*l = LevelCritical
	case LevelFatalStr:
		*l = LevelFatal
	case LevelPanicStr:
		*l = LevelPanic
	}
	return *l
}
