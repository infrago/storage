package storage

import "strings"

const (
	thumbType = "jpg"
)

type (
	Thumbnailer struct {
		Name   string
		Text   string
		Alias  []string
		Action ThumbnailFunc
	}
	ThumbnailFunc func(*File, int64, int64, int64) (string, error)

	Previewer struct {
		Name   string
		Text   string
		Alias  []string
		Action PreviewerFunc
	}
	PreviewerFunc func(*File, int64, int64, int64) (string, error)

	Info struct {
		Type string
		File string
	}
)

func (m *Module) Thumbnail(code string, width, height, position int64) (string, error) {
	file, err := decodeFile(code)
	if err != nil {
		return "", errInvalidCode
	}

	if handler, ok := m.thumbnailers[strings.ToLower(file.Type())]; ok {
		return handler.Action(file, width, height, position)
	}

	return "", errInvalidHandler
}

func (m *Module) Preview(code string, width, height, position int64) (string, error) {
	file, err := decodeFile(code)
	if err != nil {
		return "", errInvalidCode
	}

	if handler, ok := m.previewers[strings.ToLower(file.Type())]; ok {
		return handler.Action(file, width, height, position)
	}

	return "", errInvalidHandler
}
