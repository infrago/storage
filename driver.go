package storage

import (
	"io"
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

	Option struct {
		Key      string
		Root     string
		Mimetype string
		Metadata Map
		Tags     Map
		Expires  time.Time

		//for browse
		QueryString Map
	}

	Stream interface {
		io.Reader
		io.Seeker
		io.Closer
		io.ReaderAt
	}

	// Connect
	Connect interface {
		Open() error
		Health() Health
		Close() error

		Upload(string, ...Option) (string, error)
		Fetch(File, ...Option) (Stream, error)
		Download(File, ...Option) (string, error)

		Remove(File) error
		Browse(File, ...Option) (string, error)
	}
)
