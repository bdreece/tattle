package clf

import (
	"fmt"
	"time"
)

const (
	// LAYOUT provides a printf format string for the CLF format
	LAYOUT string = "%s %s %s [%s] \"%s\" %d %d\n"
	// TIME_LAYOUT provides the layout for date/time display
	TIME_LAYOUT string = "02/Jan/06:15:04 -0700"
)

// Format provides raw access to the fields of a given entry in a CLF format log
type Format struct {
	Host 		string
	Ident 		string
	AuthUser 	string
	Request 	string
	Status 		int
	Bytes 		int
}

func (f Format) Builder() Builder {
	return Builder {
		defaults: f,
	}
}

func (f Format) Format() string {
	return fmt.Sprintf(LAYOUT, f.Host, f.Ident, f.AuthUser, 
		time.Now().Format(TIME_LAYOUT), f.Request, f.Status, f.Bytes)
}
