package storage

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	. "github.com/infrago/base"
	"github.com/infrago/util"
)

type (
	Instance struct {
		Name    string
		Config  Config
		Setting Map

		connect Connect
	}
)

func (this *Instance) Hash(file string) string {
	if f, e := os.Open(file); e == nil {
		defer f.Close()
		h := sha256.New()
		if _, e := io.Copy(h, f); e == nil {
			// hex := fmt.Sprintf("%x", h.Sum(nil))
			bbb := base64.URLEncoding.EncodeToString(h.Sum(nil))
			return bbb
		}
	}
	return ""
}

func (this *Instance) File(hash string, file string, size int64) File {
	info := &filed{}

	info.base = this.Name
	info.hash = hash
	info.file = file
	info.name = path.Base(info.file)
	info.tttt = util.Extension(info.name)
	info.size = size

	info.code = encode(info)

	info.proxy = this.Config.Proxy
	info.remote = this.Config.Remote

	return info
}

// 统一返回本地缓存目录
func (this *Instance) Download(file File) (string, error) {
	hash := util.Sha256(file.Hash())
	hashPath := path.Join(hash[0:2], hash[2:4])

	full := file.Hash()
	if file.Type() != "" {
		full = fmt.Sprintf("%s.%s", file.Hash(), file.Type())
	}

	spath := path.Join(module.config.Download, hashPath)
	sfile := path.Join(spath, full)

	// //创建目录
	err := os.MkdirAll(spath, 0777)
	if err != nil {
		return "", errors.New("生成目录失败")
	}

	return sfile, nil
}
