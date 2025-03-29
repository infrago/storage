package storage

import (
	"errors"

	"github.com/infrago/infra"
	"github.com/infrago/util"
)

func (this *Module) instance(code string) (*Instance, *File, error) {
	info, err := decode(code)
	if err != nil {
		return nil, nil, errInvalidCode
	}

	base := info.Base()
	if base == "" {
		base = infra.DEFAULT
	}
	if inst, ok := this.instances[base]; ok {
		return inst, info, nil
	}
	return nil, nil, errInvalidConnection
}

func (this *Module) UploadTo(base string, orginal string, opts ...UploadOption) (*File, error) {
	if base == "" {
		base = infra.DEFAULT
	}

	inst, ok := this.instances[base]
	if ok == false {
		return nil, errInvalidConnection
	}

	opt := UploadOption{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	//默认全局的prefix
	if inst.Config.Prefix != "" && opt.Prefix == "" {
		opt.Prefix = inst.Config.Prefix
	}
	if opt.Mimetype == "" {
		//minio,s3,会自动判断
		// opt.Mimetype = infra.Mimetype(util.Extension(orginal))
	}

	return inst.connect.Upload(orginal, opt)
}

func (this *Module) Upload(orginal string, opts ...UploadOption) (*File, error) {
	//这里自动分配一个存储
	hash := util.Sha1BaseFile(orginal)
	if hash == "" {
		return nil, errors.New("hash error")
	}
	base := this.hashring.Locate(hash)

	file, err := this.UploadTo(base, orginal, opts...)
	if err != nil {
		return nil, err
	}

	return file, nil
}
func (this *Module) Fetch(code string, opts ...FetchOption) (Stream, error) {
	inst, file, err := this.instance(code)
	if err != nil {
		return nil, err
	}

	opt := FetchOption{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	return inst.connect.Fetch(file, opt)
}

func (this *Module) Download(code string, opts ...DownloadOption) (string, error) {
	inst, file, err := this.instance(code)
	if err != nil {
		return "", err
	}

	opt := DownloadOption{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	if opt.Target == "" {
		target, err := inst.downloadTarget(file)
		if err != nil {
			return "", err
		}
		opt.Target = target
	}

	return inst.connect.Download(file, opt)
}

func (this *Module) Remove(code string, opts ...RemoveOption) error {
	inst, file, err := this.instance(code)
	if err != nil {
		return err
	}

	opt := RemoveOption{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	return inst.connect.Remove(file, opt)
}

func (this *Module) Browse(code string, opts ...BrowseOption) (string, error) {
	inst, file, err := this.instance(code)
	if err != nil {
		return "", err
	}

	opt := BrowseOption{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	return inst.connect.Browse(file, opt)
}
