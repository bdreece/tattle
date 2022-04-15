package event

import "github.com/bdreece/tattle/pkg/format"

type Factory[F format.Format] struct {
	builder format.Builder[Context, F]
}

func (f Factory[F]) Create(ctx Context) string {
	return f.builder.Build(ctx).Format()
}