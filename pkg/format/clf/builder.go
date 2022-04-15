package clf

import "github.com/bdreece/tattle/pkg/context/request"

type Builder struct {
	Defaults Format
}

func (b Builder) Build(ctx request.Context) Format {
	return Format {
	}
}