package xml

import (
	"github.com/bdreece/tattle/pkg/context"
	"github.com/bdreece/tattle/pkg/format"
)

type Builder struct {
	defaults Format
}

func (b Builder) Build(ctx context.Context) format.Format {
	return Format{data: ctx.(map[string]any)}
}
