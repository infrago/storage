# storage

`storage` 是 infrago 的**模块**。

## 包定位

- 类型：模块
- 作用：对象存储模块，负责上传下载、URL 签名、元数据处理。

## 主要功能

- 对上提供统一模块接口
- 对下通过驱动接口接入具体后端
- 支持按配置切换驱动实现
- 支持按文件类型注册缩略图/预览处理器

## 快速接入

```go
import _ "github.com/infrago/storage"
```

```toml
[storage]
driver = "default"
```

```go
import "github.com/infrago/storage"
import "github.com/infrago/infra"

storage.Thumbnail(code, 320, 0, 0)

infra.Register("svg", storage.Thumbnailer{
	Action: func(file *storage.File, width, height, position int64) (string, error) {
		return "", nil
	},
})
```

## 驱动实现接口列表

以下接口由驱动实现（来自模块 `driver.go`）：

### Driver

- `Connect(*Instance) (Connection, error)`

### Stream

- `io.Reader`
- `io.Seeker`
- `io.Closer`
- `io.ReaderAt`

### Connection

- `Open() error`
- `Health() Health`
- `Close() error`
- `Upload(string, UploadOption) (*File, error)`
- `Fetch(*File, FetchOption) (Stream, error)`
- `Download(*File, DownloadOption) (string, error)`
- `Remove(*File, RemoveOption) error`
- `Browse(*File, BrowseOption) (string, error)`

## 全局配置项（所有配置键）

配置段：`[storage]`

- 未检测到配置键（请查看模块源码的 configure 逻辑）

## 说明

- `setting` 一般用于向具体驱动透传专用参数
- 多实例配置请参考模块源码中的 Config/configure 处理逻辑
