package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/infrago/infra"
	. "github.com/infrago/base"
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

	base := file.Base()
	if base == infra.DEFAULT {
		base = ""
	}

	root := filepath.Clean(module.filecfg.Download)
	rel := filepath.Clean(filepath.Join(base, file.Prefix(), name))
	if rel == "." || filepath.IsAbs(rel) || rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return "", errInvalidCode
	}

	sfile := filepath.Clean(filepath.Join(root, rel))
	if sfile != root && !strings.HasPrefix(sfile, root+string(os.PathSeparator)) {
		return "", errInvalidCode
	}
	spath := filepath.Dir(sfile)
	if err := os.MkdirAll(spath, 0o755); err != nil {
		return "", err
	}
	return sfile, nil
}
