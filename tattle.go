package tattle

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

// Logger provides configurable message prefixes for logging
type Logger struct {
	// TimeFormat contains the time layout for use in Time.Format
	TimeFormat string
	// LogPrefix contains the tag associated with log prints
	LogPrefix  string
	// WarnPrefix contains the tag associated with warning prints
	WarnPrefix string
	// ErrPrefix contains the tag associated with error prints
	ErrPrefix  string

	LogSprint  func(a ...interface{}) string
	WarnSprint func(a ...interface{}) string
	ErrSprint  func(a ...interface{}) string
}

// LoggerOpts provides optional configuration settings for Logger
type LoggerOpts struct {
	TimeFormat string
	LogColor   color.Attribute
	WarnColor  color.Attribute
	ErrColor   color.Attribute
}

// FORMAT provides the standardized log format string
const FORMAT string = "(%s) [%s] %s"

// NewLogger creates a new logger, with optional settings passed through opts.
// The default time format is time.RFC3339.
// The default log, warning, and error colors are color.BgGreen, color.BgYellow, and color.BgRed, respectively.
func NewLogger(logPrefix, warnPrefix, errPrefix string, opts *LoggerOpts) Logger {
	var (
		logColor   color.Attribute
		warnColor  color.Attribute
		errColor   color.Attribute
		timeFormat string
	)

	if opts != nil {
		logColor = opts.LogColor
		warnColor = opts.WarnColor
		errColor = opts.ErrColor
		timeFormat = opts.TimeFormat
	} else {
		logColor = color.BgGreen
		warnColor = color.BgYellow
		errColor = color.BgRed
		timeFormat = time.RFC3339
	}

	logger := Logger{
		TimeFormat: timeFormat,
		LogPrefix:  logPrefix,
		WarnPrefix: warnPrefix,
		ErrPrefix:  errPrefix,
		LogSprint:  color.New(logColor).SprintFunc(),
		WarnSprint: color.New(warnColor).SprintFunc(),
		ErrSprint:  color.New(errColor).SprintFunc(),
	}

	return logger
}

// Logf forwards format and args to fmt.Printf with log prefix and timestamp
func (l Logger) Logf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	fmt.Printf(FORMAT, l.LogSprint(l.LogPrefix), time.Now().Format(l.TimeFormat), message)
}

// Logln forwards message to fmt.Printf with log prefix and timestamp
func (l Logger) Logln(message string) {
	fmt.Printf(FORMAT+"\n", l.LogSprint(l.LogPrefix), time.Now().Format(l.TimeFormat), message)
}

// Warnf forwards format and args to fmt.Printf with warning prefix and timestamp
func (l Logger) Warnf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	fmt.Printf(FORMAT, l.WarnSprint(l.WarnPrefix), time.Now().Format(l.TimeFormat), message)
}

// Warnln forwards message to fmt.Printf with warning prefix and timestamp
func (l Logger) Warnln(message string) {
	fmt.Printf(FORMAT+"\n", l.WarnSprint(l.LogPrefix), time.Now().Format(l.TimeFormat), message)
}

// Errf forwards format and args to fmt.Printf with error prefix and timestamp
func (l Logger) Errf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	fmt.Printf(FORMAT, l.ErrSprint(l.ErrPrefix), time.Now().Format(l.TimeFormat), message)
}

// Errln forwards message to fmt.Printf with error prefix and timestamp
func (l Logger) Errln(message string) {
	fmt.Printf(FORMAT+"\n", l.ErrSprint(l.ErrPrefix), time.Now().Format(l.TimeFormat), message)
}
