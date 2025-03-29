package storage

import (
	"errors"
	"fmt"
	"image"
	"io"
	"os"
	"path"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"

	"github.com/infrago/infra"

	"github.com/disintegration/imaging"
)

func init() {
	infra.Register("image", Thumbnailer{
		Alias: []string{"jpg", "jpeg", "png", "bmp", "gif", "webp"},
		Name:  "图片处理器", Text: "图片处理器",
		Action: func(file *File, width, height, pos int64) (string, error) {
			//先获取缩略图的文件
			_, tfile, err := thumbnailing(file, width, height, pos)
			if err != nil {
				return "", err
			}

			//如果缩图已经存在，直接返回
			_, err = os.Stat(tfile)
			if err == nil {
				return tfile, nil
			}

			sfile, err := Download(file.Code())
			if err != nil {
				return "", err
			}

			sf, err := os.Open(sfile)
			if err != nil {
				return "", err
			}
			defer sf.Close()

			sf.Seek(0, 0)
			cfg, err := DecodeImageConfig(sf)
			if err != nil {
				return "", err
			}

			//流重新定位
			sf.Seek(0, 0)
			img, err := imaging.Decode(sf)
			if err != nil {
				return "", err
			}

			//计算新度和新高
			ratio := float64(cfg.Width) / float64(cfg.Height)
			newWidth, newHeight := float64(width), float64(height)
			if newWidth == 0 {
				newWidth = newHeight * ratio
			} else if newHeight == 0 {
				newHeight = newWidth / ratio
			}

			thumb := imaging.Thumbnail(img, int(newWidth), int(newHeight), imaging.NearestNeighbor)
			err = imaging.Save(thumb, tfile)
			if err != nil {
				return "", err
			}

			return tfile, nil
		},
	})

}

// 获取file的缩图路径信息
func thumbnailing(file *File, w, h, p int64) (string, string, error) {
	//使用hash的hex hash 的前4位，生成2级目录
	//共256*256个目录
	// hash := util.Sha1(file.Key())
	// hashPath := path.Join(hash[0:2], hash[2:4])

	// 待优化，这里为什么要用原始扩展名，忘记了
	ext := "jpg"
	if file.Type() != "" {
		ext = file.Type()
	}

	tpath := path.Join(ThumbnailConfig(), file.Prefix(), file.Key())
	tname := fmt.Sprintf("%d-%d-%d.%s", w, h, p, ext)
	tfile := path.Join(tpath, tname)

	// //创建目录
	err := os.MkdirAll(tpath, 0777)
	if err != nil {
		return "", "", errors.New("生成目录失败")
	}

	return tpath, tfile, nil
}

// Decode reads an image from r.
func DecodeImageConfig(r io.Reader) (image.Config, error) {

	cfg, _, err := image.DecodeConfig(r)
	if err != nil {
		return image.Config{}, err
	}
	return cfg, nil
}

// Open loads an image from file
func ImageConfig(filename string) (image.Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return image.Config{}, err
	}
	defer file.Close()
	return DecodeImageConfig(file)
}
