package logging

const (
	maxAgeDefault     = 28
	maxSizeDefault    = 500
	maxBackupsDefault = 5
)

type Options struct {
	Filename   string
	MaxAge     int
	MaxSize    int
	MaxBackups int
}

func DefaultOptions(filename string) *Options {
	return &Options{
		Filename:   filename,
		MaxAge:     maxAgeDefault,
		MaxSize:    maxSizeDefault,
		MaxBackups: maxBackupsDefault,
	}
}