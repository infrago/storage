package storage

import (
	"errors"
	"fmt"
	"os"
	"path"

	. "github.com/infrago/base"
	"github.com/infrago/infra"
)

type (
	Instance struct {
		connect Connect

		Name    string
		Config  Config
		Setting Map
	}
)

// func (this *Instance) Hash(file string) string {
// 	if f, e := os.Open(file); e == nil {
// 		defer f.Close()
// 		h := sha256.New()
// 		if _, e := io.Copy(h, f); e == nil {
// 			// hex := fmt.Sprintf("%x", h.Sum(nil))
// 			bbb := base64.URLEncoding.EncodeToString(h.Sum(nil))
// 			return bbb
// 		}
// 	}
// 	return ""
// }

func (this *Instance) File(prefix, key, tttt string, size int64) File {
	info := &filed{}

	info.base = this.Name
	info.prefix = prefix
	info.key = key
	info.tttt = tttt
	info.size = size

	info.code = encode(info)

	info.proxy = this.Config.Proxy
	info.remote = this.Config.Remote

	return info
}

// 统一返回本地缓存目录
func (this *Instance) downloadTarget(file File) (string, error) {
	name := file.Key()
	if file.Type() != "" {
		name = fmt.Sprintf("%s.%s", file.Key(), file.Type())
	}

	base := file.Base()
	if base == infra.DEFAULT {
		base = ""
	}

	sfile := path.Join(module.config.Download, file.Base(), file.Prefix(), name)
	spath := path.Dir(sfile)

	// //创建目录
	err := os.MkdirAll(spath, 0777)
	if err != nil {
		return "", errors.New("生成目录失败")
	}

	return sfile, nil
}
