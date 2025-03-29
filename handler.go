package storage

const (
	thumbType = "jpg"
)

type (
	Thumbnailer struct {
		// Name 名称
		Name string
		// Text 说明
		Text string
		// Alias 别名
		Alias []string
		// Thumbnail 缩图
		Action ThumbnailFunc
	}
	ThumbnailFunc func(*File, int64, int64, int64) (string, error)

	Previewer struct {
		// Name 名称
		Name string
		// Text 说明
		Text string
		// Alias 别名
		Alias []string
		// Preview 预览动图
		Action PreviewerFunc
	}
	PreviewerFunc func(*File, int64, int64, int64) (string, error)

	Info struct {
		Type string
		File string
	}
)

// Thumbnail 生成缩略图
func (this *Module) Thumbnail(code string, width, height, position int64) (string, error) {
	file, err := decode(code)
	if err != nil {
		return "", errInvalidCode
	}

	if handler, ok := this.thumbnailers[file.Type()]; ok {
		return handler.Action(file, width, height, position)
	}

	return "", errInvalidHandler
}

// Preview 生成预览图
func (this *Module) Preview(code string, width, height, position int64) (string, error) {
	file, err := decode(code)
	if err != nil {
		return "", errInvalidCode
	}

	if pre, ok := this.previewers[file.Type()]; ok {
		return pre.Action(file, width, height, position)
	}

	return "", errInvalidHandler
}
