package local_fs

import (
	"os"
)

func Permission(dst string) bool {
	_, err := os.Stat(dst)

	return os.IsPermission(err)
}

func CheckPath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CreatePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}
