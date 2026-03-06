package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "github.com/infrago/base"
	"github.com/infrago/infra"
)

type (
	Instance struct {
		conn Connection

		Name    string
		Config  Config
		Setting Map
	}
)

// NewFile creates a storage file metadata object for drivers.
func (i *Instance) NewFile(prefix, key, typee string, size int64) *File {
	return i.newFile(prefix, key, typee, size)
}

func (i *Instance) downloadTarget(file *File) (string, error) {
	name := file.Key()
	if file.Type() != "" {
		name = fmt.Sprintf("%s.%s", file.Key(), file.Type())
	}
	return i.prepareCacheTarget(module.filecfg.Download, filePathParts(file, name)...)
}

func (i *Instance) thumbnailTarget(file *File, name string) (string, error) {
	return i.prepareCacheTarget(module.filecfg.Thumbnail, thumbnailPathParts(file, name)...)
}

func (i *Instance) prepareCacheTarget(root string, parts ...string) (string, error) {
	target, err := cacheTarget(root, parts...)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return "", err
	}
	return target, nil
}

func cacheTarget(root string, parts ...string) (string, error) {
	root = filepath.Clean(root)
	rel := filepath.Clean(filepath.Join(parts...))
	if rel == "." || filepath.IsAbs(rel) || rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return "", errInvalidCode
	}

	target := filepath.Clean(filepath.Join(root, rel))
	if target != root && !strings.HasPrefix(target, root+string(os.PathSeparator)) {
		return "", errInvalidCode
	}
	return target, nil
}

func filePathParts(file *File, name string) []string {
	base := file.Base()
	if base == infra.DEFAULT {
		base = ""
	}
	return []string{base, file.Prefix(), name}
}

func thumbnailPathParts(file *File, name string) []string {
	base := file.Base()
	if base == infra.DEFAULT {
		base = ""
	}
	return []string{base, file.Prefix(), file.Key(), name}
}
