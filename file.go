package xs

import (
	"io"
	"os"
)

func FileExist(path string) (bool, error) {
	s, err := os.Stat(path)
	if err == nil {
		return !s.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// WARN: the function is very dangerous, use it carefully
func RemoveFile(dst string) error {
	if b, err := FileExist(dst); err != nil {
		return err
	} else if b {
		return os.Remove(dst)
	}
	return nil
}

func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
