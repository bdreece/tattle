package tattle

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	LogPrefix string
	WarnPrefix string
	ErrPrefix string
}

const FORMAT string = "[%s] (%s): %s"
var (
	green func(a ...interface{}) string = color.New(color.FgGreen).SprintFunc()
	yellow func(a ...interface{}) string = color.New(color.FgYellow).SprintFunc()
	red func(a ...interface{}) string = color.New(color.FgRed).SprintFunc()
)

func (l Logger) Logf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	fmt.Printf(FORMAT, green(l.LogPrefix), time.Now(), message)
}

func (l Logger) Logln(message string) {
	fmt.Printf(FORMAT + "\n", green(l.LogPrefix), time.Now(), message)
}

func (l Logger) Warnf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	fmt.Printf(FORMAT, yellow(l.WarnPrefix), time.Now(), message)
}

func (l Logger) Warnln(message string) {
	fmt.Printf(FORMAT + "\n", green(l.LogPrefix), time.Now(), message)
}

func (l Logger) Errf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	fmt.Printf(FORMAT, red(l.ErrPrefix), time.Now(), message)
}

func (l Logger) Errln(message string) {
	fmt.Printf(FORMAT + "\n", red(l.ErrPrefix), time.Now(), message)
}
