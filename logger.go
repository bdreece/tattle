package tattle

import (
	"io"

	"github.com/bdreece/tattle/pkg/context"
)

type Logger[C context.Context] struct {
	Dest io.Writer
	Factory context.Factory[C]
}

func (l Logger[C]) Log(ctx C) {
	l.Dest.Write([]byte(l.Factory.Create(ctx)))
}