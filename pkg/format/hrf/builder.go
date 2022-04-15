package hrf

import "github.com/bdreece/tattle/pkg/context"

type Builder[C context.Context] struct {
	Defaults Format
}

func (b Builder[C]) Build(ctx C) Format {
	return Format{}
}