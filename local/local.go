package local

import (
	"errors"
	"net/url"
	"os"

	"github.com/aeppert/stow"
)

// ConfigKeys are the supported configuration items for
// local storage.
const (
	ConfigKeyPath = "path"
)

// Kind is the kind of Location this package provides.
const Kind = "local"

const (
	paramTypeValue = "item"
)

func init() {
	makefn := func(config stow.Config) (stow.Location, error) {
		path, ok := config.Config(ConfigKeyPath)
		if !ok {
			return nil, errors.New("missing path config")
		}
		info, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		if !info.IsDir() {
			return nil, errors.New("path must be directory")
		}
		return &location{
			config: config,
		}, nil
	}
	kindfn := func(u *url.URL) bool {
		return u.Scheme == "file"
	}
	stow.Register(Kind, makefn, kindfn)
}
