package storage

import (
	"io"
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

		Upload(string, UploadOption) (string, error)
		Fetch(File, FetchOption) (Stream, error)
		Download(File, DownloadOption) (string, error)

		Remove(File, RemoveOption) error
		Browse(File, BrowseOption) (string, error)
	}
)
