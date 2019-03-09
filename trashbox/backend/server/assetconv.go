package server

type (
	Format struct {
		GZIP bool
	}
	Converter struct {
	}
)

func (c Converter) isSupport(contentType, format string) bool {
	if format == "" {
		return true
	}
	switch contentType {
	case "image/jpeg":
		fallthrough
	case "image/png":
		return format == "icon" || format == "thumbnail"
	}
	return false
}
