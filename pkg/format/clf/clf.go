package clf

import (
	"fmt"
	"net/http"
	"time"
)

const (
	FORMAT string = "%s %s %s [%s] \"%s\" %d %d\n"
	TIME_LAYOUT string = "02/Jan/06:15:04 -0700"
)

type ClfFormat struct {
	Host 		string
	Ident 		string
	AuthUser 	string
	Request 	string
	Status 		int
	Bytes 		int
}

type ClfBuilder struct {}

func (b ClfBuilder) FromRequest(r *http.Request) *ClfFormat {
	return &ClfFormat {
		Host: r.Host,
		Ident: "-",
		AuthUser: "-",
		Request: fmt.Sprintf("%s %s %s", r.Method, r.URL, r.Proto),
	}
}

func (l *ClfFormat) WithStatus(status int) *ClfFormat {
	l.Status = status
	return l
}

func (l *ClfFormat) WithBytes(bytes int) *ClfFormat {
	l.Bytes = bytes
	return l
}

func (l ClfFormat) Marshal() string {
	return fmt.Sprintf(FORMAT, l.Host, l.Ident, l.AuthUser, 
		time.Now().Format(TIME_LAYOUT), l.Request, l.Status, l.Bytes)
}