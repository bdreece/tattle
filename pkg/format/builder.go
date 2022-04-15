package format

import "github.com/bdreece/tattle/pkg/context"

type Builder[C context.Context, F Format] interface {
	Build(ctx C) F
}