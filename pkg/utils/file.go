package utils

import (
	"gitee.com/KongchengPro/GoBuilder/internal/app"
	"os"
	"path/filepath"
)

// MustMkDirAll behaves the same as `os.Mkdir`, but it will panic an error instead of return it.
func MustMkDirAll(basePath, dirName string) {
	err := os.MkdirAll(filepath.Join(basePath, dirName), app.DefaultPerm)
	if err != nil {
		panic(err)
	}
}

// MustMkFile create a file then close it, if there is an error, it will panic it
func MustMkFile(dirPath, fileName string) {
	f, err := os.Create(filepath.Join(dirPath, fileName))
	defer func() {
		newErr := f.Close()
		if newErr != nil {
			panic(newErr)
		}
	}()
	if err != nil {
		panic(err)
	}
}

// IsExist judge whether the file/directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	// `err == nil` means the file/directory is exists,
	// but if `err != nil` does not means the file/directory is not exists,
	// so call `os.IsExist` to judge.
	return err == nil || os.IsExist(err)
}
