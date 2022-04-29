package request

import (
	"io"
	"net/http"

	"github.com/bdreece/tattle"
	"github.com/bdreece/tattle/pkg/format"
)

func NewLogger(dest io.Writer, builder format.Builder[Context]) tattle.Logger[Context] {
	return tattle.NewLogger(dest, builder)
}

func LogAndRespond(l tattle.Logger[Context], r *http.Request, w http.ResponseWriter, body []byte) error {
	n, err := w.Write(body)
	if err != nil {
		return err
	}
	l.Log(Context {
		Request: r,
		Status: http.StatusOK,
		Bytes: n,
	})
	return nil
}