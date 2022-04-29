package event

import (
	"io"

	"github.com/bdreece/tattle"
	"github.com/bdreece/tattle/pkg/format"
)

func NewLogger(dest io.Writer, builder format.Builder[Context]) tattle.Logger[Context] {
	return tattle.NewLogger(dest, builder)
}