package tattle

import (
	"io"

	_ "github.com/bdreece/tattle/pkg/context"
	_ "github.com/bdreece/tattle/pkg/format"
)

type Logger[C Context] struct {
	dest io.Writer
	factory Factory[C]
}

func (l Logger[C]) Log(ctx C) {
	l.dest.Write([]byte(l.factory.Create(ctx)))
}