package format

// Format provides an interface for log formats (e.g. CLF, ELF, JSON, etc.)
type Format interface {
	Format() string
}