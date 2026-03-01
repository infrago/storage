# storage

`storage` 是 infrago 的模块包。

## 安装

```bash
go get github.com/infrago/storage@latest
```

## 最小接入

```go
package main

import (
    _ "github.com/infrago/storage"
    "github.com/infrago/infra"
)

func main() {
    infra.Run()
}
```

## 配置示例

```toml
[storage]
driver = "default"
```

## 公开 API（摘自源码）

- `func (m *Module) UploadTo(base string, original string, opts ...UploadOption) (*File, error)`
- `func (m *Module) Upload(original string, opts ...UploadOption) (*File, error)`
- `func (m *Module) Fetch(code string, opts ...FetchOption) (Stream, error)`
- `func (m *Module) Download(code string, opts ...DownloadOption) (string, error)`
- `func (m *Module) Remove(code string, opts ...RemoveOption) error`
- `func (m *Module) Browse(code string, opts ...BrowseOption) (string, error)`
- `func (i *Instance) NewFile(prefix, key, typee string, size int64) *File`
- `func (f *File) Base() string   { return f.base }`
- `func (f *File) Prefix() string { return f.prefix }`
- `func (f *File) Key() string    { return f.key }`
- `func (f *File) Type() string   { return f.typee }`
- `func (f *File) Size() int64    { return f.size }`
- `func (f *File) Code() string   { return f.code }`
- `func (f *File) Proxy() bool    { return f.proxy }`
- `func (f *File) Remote() bool   { return f.remote }`
- `func (f *File) File() string`
- `func (f *File) Name() string`
- `func Decode(code string) (*File, error)`
- `func (d *defaultDriver) Connect(instance *Instance) (Connection, error)`
- `func (c *defaultConnection) Open() error { return nil }`
- `func (c *defaultConnection) Health() Health`
- `func (c *defaultConnection) Close() error { return nil }`
- `func (c *defaultConnection) Upload(original string, opt UploadOption) (*File, error)`
- `func (c *defaultConnection) Fetch(file *File, opt FetchOption) (Stream, error)`
- `func (c *defaultConnection) Download(file *File, opt DownloadOption) (string, error)`
- `func (c *defaultConnection) Remove(file *File, _ RemoveOption) error`
- `func (c *defaultConnection) Browse(file *File, _ BrowseOption) (string, error)`
- `func Upload(from Any, opts ...UploadOption) (*File, error)`
- `func UploadTo(base string, from Any, opts ...UploadOption) (*File, error)`
- `func Fetch(code string, opts ...FetchOption) (Stream, error)`
- `func Download(code string, opts ...DownloadOption) (string, error)`
- `func Remove(code string, opts ...RemoveOption) error`
- `func Browse(code string, opts ...BrowseOption) (string, error)`
- `func ThumbnailConfig() string { return module.filecfg.Thumbnail }`
- `func PreviewConfig() string   { return module.filecfg.Preview }`
- `func SaltConfig() string      { return module.filecfg.Salt }`
- `func (r *rangeFileReader) Read(p []byte) (int, error)`
- `func (r *rangeFileReader) Seek(offset int64, whence int) (int64, error)`
- `func (r *rangeFileReader) ReadAt(p []byte, off int64) (int, error)`
- `func (r *rangeFileReader) Close() error`

## 排错

- 模块未运行：确认空导入已存在
- driver 无效：确认驱动包已引入
- 配置不生效：检查配置段名是否为 `[storage]`
