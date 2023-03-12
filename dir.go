package xs

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func DirExist(path string) (bool, error) {
	s, err := os.Stat(path)
	if err == nil {
		return s.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func PresistDir(path string) error {
	if exist, err := DirExist(path); err != nil {
		return err
	} else if exist {
		return nil
	}
	return os.MkdirAll(path, os.ModePerm)
}

func JoinURL(base string, paths ...string) string {
	p := path.Join(paths...)
	return fmt.Sprintf("%s/%s", strings.TrimRight(base, "/"), strings.TrimLeft(p, "/"))
}

func IsEmptyDir(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()
	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}
