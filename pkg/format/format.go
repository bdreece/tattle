package format

// LogFormat provides an interface for log formats (e.g. CLF, ELF, JSON, etc.)
type LogFormat interface {
	SetField(key string, val any) error
	Format() string
}