package request

import "net/http"

type Context struct {
	Request *http.Request
	Status	int
	Bytes	int
}