package storage

import (
	"time"

	. "github.com/infrago/base"
)

func (this *Module) UploadTo(base string, path string, metadata Map) (File, Files, error) {
	if inst, ok := this.instances[base]; ok {
		return inst.connect.Upload(path, metadata)
	}
	return nil, nil, errInvalidConnection
}

func (this *Module) Upload(path string, metadata Map) (File, Files, error) {
	//这里自动分配一个存储
	base := this.hashring.Locate(path)
	return this.UploadTo(base, path, metadata)
}

func (this *Module) Download(code string) (string, error) {
	info, err := decode(code)

	if err != nil {
		return "", errInvalidCode
	}

	if inst, ok := this.instances[info.Base()]; ok {
		file, err := inst.connect.Download(info)
		if err != nil {
			return "", err
		}

		// info.file = file
		// info.name = path.Base(info.file)

		return file, nil

	}
	return "", errInvalidConnection
}

func (this *Module) Remove(code string) error {
	info, err := decode(code)
	if err != nil {
		return errInvalidCode
	}

	if inst, ok := this.instances[info.Base()]; ok {
		return inst.connect.Remove(info)
	}
	return errInvalidConnection
}

func (this *Module) Browse(code string, query Map, expires ...time.Duration) (string, error) {
	exp := time.Duration(0)
	if len(expires) > 0 {
		exp = expires[0]
	}

	info, err := decode(code)
	if err != nil {
		return "", errInvalidCode
	}

	if inst, ok := this.instances[info.Base()]; ok {
		return inst.connect.Browse(info, query, exp)
	}

	return "", errInvalidConnection
}
