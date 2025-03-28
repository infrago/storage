package storage

import (
	"errors"
	"io"

	"github.com/infrago/util"
)

func (this *Module) UploadTo(base string, orginal string, opts ...Option) (string, error) {
	if inst, ok := this.instances[base]; ok {
		return inst.connect.Upload(orginal, opts...)
	}
	return "", errInvalidConnection
}

func (this *Module) Upload(orginal string, opts ...Option) (string, error) {
	//这里自动分配一个存储
	hash := util.Sha1BaseFile(orginal)
	if hash == "" {
		return "", errors.New("hash error 123")
	}
	base := this.hashring.Locate(hash)
	return this.UploadTo(base, orginal, opts...)
}
func (this *Module) Fetch(code string, opts ...Option) (io.Reader, error) {
	info, err := decode(code)

	if err != nil {
		return nil, errInvalidCode
	}

	if inst, ok := this.instances[info.Base()]; ok {
		return inst.connect.Fetch(info, opts...)
	}

	return nil, errInvalidConnection
}

func (this *Module) Download(code string, opts ...Option) (string, error) {
	info, err := decode(code)

	if err != nil {
		return "", errInvalidCode
	}

	if inst, ok := this.instances[info.Base()]; ok {
		return inst.connect.Download(info)
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

func (this *Module) Browse(code string, opts ...Option) (string, error) {
	info, err := decode(code)
	if err != nil {
		return "", errInvalidCode
	}

	if inst, ok := this.instances[info.Base()]; ok {
		return inst.connect.Browse(info, opts...)
	}

	return "", errInvalidConnection
}
