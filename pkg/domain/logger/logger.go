/*
Copyright 2018-2020 The Nori Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logger

import "time"

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

//go:generate mockgen -destination=../mocks/logger/formatter.go -package=mocks github.com/nori-io/common/v3/pkg/domain/logger Formatter
type Formatter interface {
	Format(e Entry, field ...Field) ([]byte, error)
}

//go:generate mockgen -destination=../mocks/logger/logger.go -package=mocks github.com/nori-io/common/v3/pkg/domain/logger Logger
type Logger interface {
	FieldLogger
	AddHook(hook Hook)
	With(fields ...Field) Logger
}

//go:generate mockgen -destination=../mocks/logger/field_logger.go -package=mocks github.com/nori-io/common/v3/pkg/domain/logger FieldLogger
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
