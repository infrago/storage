package storage

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/infrago/infra"
)

type testDownloadConnection struct {
	path string
}

func (c *testDownloadConnection) Open() error                                    { return nil }
func (c *testDownloadConnection) Health() Health                                 { return Health{} }
func (c *testDownloadConnection) Close() error                                   { return nil }
func (c *testDownloadConnection) Upload(string, UploadOption) (*File, error)     { return nil, nil }
func (c *testDownloadConnection) Fetch(*File, FetchOption) (Stream, error)       { return nil, nil }
func (c *testDownloadConnection) Download(*File, DownloadOption) (string, error) { return c.path, nil }
func (c *testDownloadConnection) Remove(*File, RemoveOption) error               { return nil }
func (c *testDownloadConnection) Browse(*File, BrowseOption) (string, error)     { return c.path, nil }

func TestModuleRegisterThumbnailerAliases(t *testing.T) {
	m := &Module{
		thumbnailers: make(map[string]Thumbnailer),
		previewers:   make(map[string]Previewer),
	}

	m.Register("IMAGE", Thumbnailer{Alias: []string{"JPG", " jpg "}})

	if _, ok := m.thumbnailers["image"]; !ok {
		t.Fatalf("expected primary thumbnailer alias to be registered")
	}
	if _, ok := m.thumbnailers["jpg"]; !ok {
		t.Fatalf("expected normalized thumbnailer alias to be registered")
	}
}

func TestThumbnailGeneratesImage(t *testing.T) {
	tempDir := t.TempDir()
	source := filepath.Join(tempDir, "source.png")
	if err := writeSamplePNG(source, 120, 60); err != nil {
		t.Fatalf("write source image: %v", err)
	}

	oldThumb := module.filecfg.Thumbnail
	oldInstances := module.instances
	defer func() {
		module.filecfg.Thumbnail = oldThumb
		module.instances = oldInstances
	}()

	module.filecfg.Thumbnail = filepath.Join(tempDir, "thumb")
	module.instances = map[string]*Instance{
		infra.DEFAULT: &Instance{
			Name: infra.DEFAULT,
			conn: &testDownloadConnection{path: source},
		},
	}

	inst := &Instance{Name: infra.DEFAULT}
	file := inst.NewFile("aa/bb", "sample", "png", 0)

	got, err := Thumbnail(file.Code(), 20, 0, 0)
	if err != nil {
		t.Fatalf("generate thumbnail: %v", err)
	}

	if _, err := os.Stat(got); err != nil {
		t.Fatalf("thumbnail not written: %v", err)
	}

	cfg, err := ImageConfig(got)
	if err != nil {
		t.Fatalf("decode thumbnail config: %v", err)
	}
	if cfg.Width != 20 || cfg.Height != 10 {
		t.Fatalf("unexpected thumbnail size: got %dx%d", cfg.Width, cfg.Height)
	}

	got2, err := Thumbnail(file.Code(), 20, 0, 0)
	if err != nil {
		t.Fatalf("reuse thumbnail cache: %v", err)
	}
	if got2 != got {
		t.Fatalf("expected cached thumbnail path, got %q want %q", got2, got)
	}
}

func TestThumbnailTypeFallsBackForWebP(t *testing.T) {
	if got := thumbnailType("webp"); got != thumbType {
		t.Fatalf("unexpected thumbnail type for webp: %q", got)
	}
}

func writeSamplePNG(filename string, width, height int) error {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	fill := color.RGBA{R: 0x3c, G: 0x84, B: 0xc6, A: 0xff}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, fill)
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
