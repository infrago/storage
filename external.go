package store

import (
	"errors"
	"fmt"
	"time"

	. "github.com/infrago/base"
)

func ThumbnailConfig() string {
	return module.config.Thumbnail
}
func PreviewConfig() string {
	return module.config.Preview
}
func SaltConfig() string {
	return module.config.Salt
}

func Upload(from Any, metadatas ...Map) (File, Files, error) {
	path := ""
	switch vv := from.(type) {
	case string:
		path = vv
	case Map:
		if file, ok := vv["file"].(string); ok {
			path = file
		} else {
			return nil, nil, errors.New("invalid target")
		}
	default:
		path = fmt.Sprintf("%v", vv)
	}

	var metadata Map
	if len(metadatas) > 0 {
		metadata = metadatas[0]
	}

	return module.Upload(path, metadata)
}

func UploadFile(path Any, metadatas ...Map) (File, error) {
	file, _, err := Upload(path, metadatas...)
	return file, err
}
func UploadPath(path Any, metadatas ...Map) (Files, error) {
	_, files, err := Upload(path, metadatas...)
	return files, err
}

func UploadTo(base string, from Any, metadatas ...Map) (File, Files, error) {
	path := ""
	switch vv := from.(type) {
	case string:
		path = vv
	case Map:
		if file, ok := vv["file"].(string); ok {
			path = file
		} else {
			return nil, nil, errors.New("invalid target")
		}
	default:
		path = fmt.Sprintf("%v", vv)
	}

	var metadata Map
	if len(metadatas) > 0 {
		metadata = metadatas[0]
	}

	return module.UploadTo(base, path, metadata)
}

func UploadFileTo(base string, path Any, metadatas ...Map) (File, error) {
	file, _, err := UploadTo(base, path, metadatas...)
	return file, err
}
func UploadPathTo(base string, path Any, metadatas ...Map) (Files, error) {
	_, files, err := UploadTo(base, path, metadatas...)
	return files, err
}

func Download(code string) (string, error) {
	return module.Download(code)
}
func Remove(code string) error {
	return module.Remove(code)
}

func Browse(code string, query Map, expires ...time.Duration) (string, error) {
	return module.Browse(code, query, expires...)
}

func Thumbnail(code string, width, height, pos int64) (string, error) {
	return module.Thumbnail(code, width, height, pos)
}

func Preview(code string, width, height, second int64) (string, error) {
	return module.Preview(code, width, height, second)
}
