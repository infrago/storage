package storage

import (
	"time"

	. "github.com/infrago/base"
)

type (
	// Driver
	Driver interface {
		Connect(*Instance) (Connect, error)
	}

	// Health
	Health struct {
		Workload int64
	}

	// Connect
	Connect interface {
		Open() error
		Health() Health
		Close() error

		Upload(path string, metadata Map) (File, Files, error)
		Download(file File) (string, error)
		Browse(file File, query Map, expires time.Duration) (string, error)
		Remove(file File) error
	}
)
