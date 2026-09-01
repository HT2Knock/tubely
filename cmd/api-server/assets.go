package main

import (
	"os"
)

func (cfg config) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.AssetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.AssetsRoot, 0o750)
	}
	return nil
}
