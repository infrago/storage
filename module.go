package storage

import (
	"sync"

	. "github.com/infrago/base"
	"github.com/infrago/infra"
	"github.com/infrago/util"
)

func init() {
	infra.Mount(module)
}

var (
	module = &Module{
		config: fileConfig{
			"store/download", "store/thumbnail", "store/preview", infra.INFRA,
		},
		configs: make(Configs, 0),
		drivers: make(map[string]Driver, 0),

		thumbnailers: make(map[string]Thumbnailer, 0),
		previewers:   make(map[string]Previewer, 0),

		instances: make(map[string]*Instance, 0),
	}
)

type (
	Module struct {
		mutex sync.Mutex

		connected, initialized, launched bool

		config       fileConfig
		configs      Configs
		drivers      map[string]Driver
		thumbnailers map[string]Thumbnailer
		previewers   map[string]Previewer

		instances map[string]*Instance

		weights  map[string]int
		hashring *util.HashRing
	}

	//这个是存储模块的全局配置
	fileConfig struct {
		Download  string
		Thumbnail string
		Preview   string
		Salt      string
	}

	Configs map[string]Config
	Config  struct {
		Driver  string
		Weight  int
		Prefix  string
		Proxy   bool
		Remote  bool
		Setting Map
	}
)

func (module *Module) Driver(name string, driver Driver) {
	module.mutex.Lock()
	defer module.mutex.Unlock()

	if infra.Override() {
		module.drivers[name] = driver
	} else {
		if module.drivers[name] == nil {
			module.drivers[name] = driver
		}
	}
}

func (this *Module) Config(name string, config Config) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if name == "" {
		name = infra.DEFAULT
	}

	if infra.Override() {
		this.configs[name] = config
	} else {
		if _, ok := this.configs[name]; ok == false {
			this.configs[name] = config
		}
	}
}
func (this *Module) Configs(config Configs) {
	for key, val := range config {
		this.Config(key, val)
	}
}

func (this *Module) Thumbnailer(name string, config Thumbnailer) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	alias := make([]string, 0)
	if name != "" {
		alias = append(alias, name)
	}
	if config.Alias != nil {
		alias = append(alias, config.Alias...)
	}

	for _, key := range alias {
		if infra.Override() {
			this.thumbnailers[key] = config
		} else {
			if _, ok := this.thumbnailers[key]; ok == false {
				this.thumbnailers[key] = config
			}
		}
	}
}

func (this *Module) Previewer(name string, config Previewer) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	alias := make([]string, 0)
	if name != "" {
		alias = append(alias, name)
	}
	if config.Alias != nil {
		alias = append(alias, config.Alias...)
	}

	for _, key := range alias {
		if infra.Override() {
			this.previewers[key] = config
		} else {
			if _, ok := this.previewers[key]; ok == false {
				this.previewers[key] = config
			}
		}
	}
}
