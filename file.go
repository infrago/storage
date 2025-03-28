package storage

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	. "github.com/infrago/base"

	"github.com/infrago/infra"
	"github.com/infrago/util"
)

type (
	File struct {
		base   string
		prefix string
		key    string
		tttt   string
		size   int64

		code   string
		proxy  bool
		remote bool

		// file string
		// name string
		// code string

	}

	// File interface {
	// 	Base() string
	// 	Prefix() string
	// 	Key() string
	// 	Type() string
	// 	Size() int64

	// 	// File() string
	// 	// Name() string

	// 	Code() string
	// 	Proxy() bool
	// 	Remote() bool
	// }
	Files []File
)

func (sf *File) Base() string {
	return sf.base
}
func (sf *File) Prefix() string {
	return sf.prefix
}
func (sf *File) Key() string {
	return sf.key
}
func (sf *File) Type() string {
	return sf.tttt
}
func (sf *File) Size() int64 {
	return sf.size
}

func (sf *File) Code() string {
	return sf.code
}

func (sf *File) Proxy() bool {
	return sf.proxy
}

func (sf *File) Remote() bool {
	return sf.remote
}

func (sf *File) File() string {
	return fmt.Sprintf("%s.%s", path.Join(sf.prefix, sf.key), sf.tttt)
}
func (sf *File) Name() string {
	return fmt.Sprintf("%s.%s", path.Base(sf.key), sf.tttt)
}

// func NewFile(base, hash, filepath string, size int64) File {
// 	file := &filed{}

// 	file.base = base
// 	file.hash = hash
// 	file.path = filepath
// 	file.name = path.Base(file.path)
// 	file.tttt = util.Extension(file.name)
// 	file.size = size

// 	file.code = encode(file)

// 	return file
// }

// 文件编解码
// fileConfig可以设置加解密方式
func encode(file *File) string {
	base := file.Base()
	if base == infra.DEFAULT {
		base = ""
	}
	code := fmt.Sprintf("%s\t%s\t%s\t%s\t%d", base, file.Prefix(), file.Key(), file.Type(), file.Size())
	if val, err := infra.EncryptTEXT(code); err == nil {
		return val
	}
	return ""
}

func decode(code string) (*File, error) {
	val, err := infra.DecryptTEXT(code)
	if err != nil {
		return nil, err
	}

	args := strings.Split(fmt.Sprintf("%v", val), "\t")
	if len(args) != 5 {
		return nil, errInvalidCode
	}

	info := &File{}
	info.code = code
	info.base = args[0]
	info.prefix = args[1]
	info.key = args[2]
	info.tttt = args[3]
	if vv, err := strconv.ParseInt(args[4], 10, 64); err == nil {
		info.size = vv
	}

	//加上状态信息
	if cfg, ok := module.configs[info.base]; ok {
		info.proxy = cfg.Proxy
		info.remote = cfg.Remote
	}

	return info, nil
}

func StatFile(file string) (Map, error) {
	stat, err := os.Stat(file)
	if err != nil {
		return nil, err
	}

	hash := util.Sha1BaseFile(file)
	if hash == "" {
		return nil, errors.New("hash error")
	}
	filename := stat.Name()
	extension := util.Extension(file)
	mimetype := infra.Mimetype(extension)
	length := stat.Size()

	return Map{
		"hash": hash,
		"name": filename,
		"type": extension,
		"mime": mimetype,
		"size": length,
		"file": file,
	}, nil
}

func Decode(code string) (*File, error) {
	return decode(code)
}
