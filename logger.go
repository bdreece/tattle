package tattle

import (
	"io"

	"github.com/bdreece/tattle/pkg/context"
	"github.com/bdreece/tattle/pkg/format"
)

type Logger[C context.Context] struct {
	dest io.Writer
	builder format.Builder[C] 
}

func NewLogger[C context.Context](dest io.Writer, builder format.Builder[C]) Logger[C] {
	return Logger[C] {
		dest,
		builder,
	}
}

func (l Logger[C]) Log(ctx C) {
	format := l.builder.Build(ctx)
	log := format.Format()
	l.dest.Write([]byte(log))
}
