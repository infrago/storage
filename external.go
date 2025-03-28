package storage

import (
	"errors"
	"fmt"
	"io"

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

func Upload(from Any, opts ...Option) (string, error) {
	path := ""
	switch vv := from.(type) {
	case string:
		path = vv
	case Map:
		if file, ok := vv["file"].(string); ok {
			path = file
		} else {
			return "", errors.New("invalid target")
		}
	default:
		path = fmt.Sprintf("%v", vv)
	}

	return module.Upload(path, opts...)
}

// func UploadFile(path Any, opts ...Option) (File, error) {
// 	file, _, err := Upload(path, metadatas...)
// 	return file, err
// }
// func UploadPath(path Any, metadatas ...Map) (Files, error) {
// 	_, files, err := Upload(path, metadatas...)
// 	return files, err
// }

func UploadTo(base string, from Any, opts ...Option) (string, error) {
	path := ""
	switch vv := from.(type) {
	case string:
		path = vv
	case Map:
		if file, ok := vv["file"].(string); ok {
			path = file
		} else {
			return "", errors.New("invalid target")
		}
	default:
		path = fmt.Sprintf("%v", vv)
	}

	return module.UploadTo(base, path, opts...)
}

// func UploadFileTo(base string, path Any, metadatas ...Map) (File, error) {
// 	file, _, err := UploadTo(base, path, metadatas...)
// 	return file, err
// }
// func UploadPathTo(base string, path Any, metadatas ...Map) (Files, error) {
// 	_, files, err := UploadTo(base, path, metadatas...)
// 	return files, err
// }

func Fetch(code string, opts ...Option) (io.Reader, error) {
	return module.Fetch(code)
}

func Download(code string, opts ...Option) (string, error) {
	return module.Download(code)
}
func Remove(code string) error {
	return module.Remove(code)
}

func Browse(code string, opts ...Option) (string, error) {
	return module.Browse(code, opts...)
}

func Thumbnail(code string, width, height, pos int64) (string, error) {
	return module.Thumbnail(code, width, height, pos)
}

func Preview(code string, width, height int64) (string, error) {
	return module.Preview(code, width, height, 0)
}
