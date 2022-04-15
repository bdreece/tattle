package elf

import "github.com/bdreece/tattle/pkg/context/event"

type Builder struct {
	Defaults Format
}

func (b Builder) Build(ctx event.Context) Format {
	return Format{}
}