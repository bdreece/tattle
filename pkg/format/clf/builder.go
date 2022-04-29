package clf

import "github.com/bdreece/tattle/pkg/context/request"

type Builder struct {
	defaults Format
}

func (b Builder) Build(ctx *request.Context) (f Format) {
	f = b.defaults
	if ctx != nil {
		f.Host = ctx.Request.Host
		f.Request = ctx.Request.Method + " " + ctx.Request.RequestURI
		f.Status = ctx.Status
		f.Bytes = ctx.Bytes
	}
	return
}

