package request

import _ "github.com/bdreece/tattle/pkg/format"

type Factory[F Format] struct {
	builder Builder[Context, F]
}

func (f Factory) Create(ctx Context) string {
	return f.builder.Build(ctx).Format()
}