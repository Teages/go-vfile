package vfile

import (
	"embed"
	"io/fs"
	"os"
	"path"
)

var tempPath string

func init() {
	tempPath, _ = os.MkdirTemp("", "go-vfile-*")
	os.TempDir()
}

func Join(filePath string, data []byte) error {
	dir := GetPath(path.Dir(filePath))
	_, err := os.Stat(dir)
	if err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(dir, 0750)
		}
	}
	return os.WriteFile(GetPath(filePath), data, 0644)
}

func JoinAll(filePath string, f embed.FS) error {
	return fs.WalkDir(f, ".", func(p string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			fileData, err := f.ReadFile(p)
			if err != nil {
				return err
			}
			return Join(path.Join(filePath, p), fileData)
		}
		return nil
	})
}

func Remove(filePath string) error {
	fileInfo, err := os.Stat(GetPath(filePath))
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		os.RemoveAll(GetPath(filePath))
	}
	return nil
}

func GetPath(filePath string) string {
	return path.Join(tempPath, filePath)
}

func Close() {
	os.RemoveAll(tempPath)
}
