package hrf

import (
	"fmt"
	"time"
)

const FORMAT string = "(%s) %s\n"

type Format struct {
	Time time.Time
	Message string
}

func (f Format) Format() string {
	return fmt.Sprintf(FORMAT, f.Time, f.Message)
}