package storage

import (
	"fmt"
	"image"
	"io"
	"os"
	"strings"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"
	"golang.org/x/sync/singleflight"

	"github.com/disintegration/imaging"
	"github.com/infrago/infra"
)

var (
	thumbnailGroup singleflight.Group

	thumbnailSaveTypes = map[string]struct{}{
		"bmp":  {},
		"gif":  {},
		"jpg":  {},
		"jpeg": {},
		"png":  {},
		"tif":  {},
		"tiff": {},
	}
)

func init() {
	infra.Register("image", Thumbnailer{
		Alias: []string{"jpg", "jpeg", "png", "bmp", "gif", "webp"},
		Name:  "图片处理器",
		Text:  "图片处理器",
		Action: func(file *File, width, height, pos int64) (string, error) {
			return generateImageThumbnail(file, width, height, pos)
		},
	})
}

func generateImageThumbnail(file *File, width, height, pos int64) (string, error) {
	inst, err := thumbnailInstance(file)
	if err != nil {
		return "", err
	}

	name := fmt.Sprintf("%d-%d-%d.%s", width, height, pos, thumbnailType(file.Type()))
	target, err := inst.thumbnailTarget(file, name)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(target); err == nil {
		return target, nil
	} else if !os.IsNotExist(err) {
		return "", err
	}

	value, err, _ := thumbnailGroup.Do(target, func() (any, error) {
		if _, err := os.Stat(target); err == nil {
			return target, nil
		} else if !os.IsNotExist(err) {
			return "", err
		}

		source, err := Download(file.Code())
		if err != nil {
			return "", err
		}

		img, err := imaging.Open(source, imaging.AutoOrientation(true))
		if err != nil {
			return "", err
		}

		tw, th := thumbnailSize(img.Bounds(), width, height)
		thumb := imaging.Thumbnail(img, tw, th, imaging.NearestNeighbor)
		if err := imaging.Save(thumb, target); err != nil {
			return "", err
		}
		return target, nil
	})
	if err != nil {
		return "", err
	}
	return value.(string), nil
}

func thumbnailInstance(file *File) (*Instance, error) {
	base := file.Base()
	if base == "" {
		base = infra.DEFAULT
	}
	inst, ok := module.instances[base]
	if !ok {
		return nil, errInvalidConnection
	}
	return inst, nil
}

func thumbnailType(typee string) string {
	typee = strings.ToLower(strings.TrimSpace(typee))
	if _, ok := thumbnailSaveTypes[typee]; ok {
		return typee
	}
	return thumbType
}

func thumbnailSize(bounds image.Rectangle, width, height int64) (int, int) {
	ow := bounds.Dx()
	oh := bounds.Dy()
	if ow <= 0 || oh <= 0 {
		return 1, 1
	}

	switch {
	case width <= 0 && height <= 0:
		width, height = int64(ow), int64(oh)
	case width <= 0:
		width = maxInt64(1, int64(float64(height)*float64(ow)/float64(oh)))
	case height <= 0:
		height = maxInt64(1, int64(float64(width)*float64(oh)/float64(ow)))
	}

	return maxInt(1, int(width)), maxInt(1, int(height))
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func DecodeImageConfig(r io.Reader) (image.Config, error) {
	cfg, _, err := image.DecodeConfig(r)
	if err != nil {
		return image.Config{}, err
	}
	return cfg, nil
}

func ImageConfig(filename string) (image.Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return image.Config{}, err
	}
	defer file.Close()
	return DecodeImageConfig(file)
}
