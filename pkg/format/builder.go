package format

import _ "github.com/bdreece/tattle/pkg/context"

type Builder[C Context, F Format] interface {
	Build(ctx C) F
}