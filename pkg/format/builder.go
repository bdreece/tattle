package format

import "github.com/bdreece/tattle/pkg/context"

type Builder[C context.Context] interface {
	Build(ctx C) Format
}
