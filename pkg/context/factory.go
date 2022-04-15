package context

type Factory[C Context] interface {
	Create(ctx C) string
}