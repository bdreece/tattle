package clf

import (
	"fmt"
	"time"
)

const (
	FORMAT string = "%s %s %s [%s] \"%s\" %d %d\n"
	TIME_LAYOUT string = "02/Jan/06:15:04 -0700"
)

type Format struct {
	Host 		string
	Ident 		string
	AuthUser 	string
	Request 	string
	Status 		int
	Bytes 		int
}

func (f Format) Format() string {
	return fmt.Sprintf(FORMAT, f.Host, f.Ident, f.AuthUser, 
		time.Now().Format(TIME_LAYOUT), f.Request, f.Status, f.Bytes)
}