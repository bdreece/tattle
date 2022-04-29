package message

import (
	"io"

	"github.com/bdreece/tattle"
	"github.com/bdreece/tattle/pkg/format"
)

func NewLogger(dest io.Writer, builder format.Builder[Context]) tattle.Logger[Context] {
	return tattle.NewLogger[Context](dest, builder)
}

func Log(l tattle.Logger, message string) {
	l.Log(Context {
		Message: message,
	})
}