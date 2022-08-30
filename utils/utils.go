package utils

import (
	"path/filepath"
)

func JoinURL(paths ...string) string {
	return filepath.ToSlash(filepath.Join(paths...))
}
