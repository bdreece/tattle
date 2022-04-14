package builder

// LogBuilder provides an interface for building logs
type LogBuilder[F LogFormat] interface {
	Build() F
}