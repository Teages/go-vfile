package vfile

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

var tempPath string

func init() {
	tempPath, _ = os.MkdirTemp("", "go-vfile-*")
	os.TempDir()
}

func Join(fileDir string, data []byte) error {
	dir := filepath.Dir(GetPath(fileDir))
	_, err := os.Stat(dir)
	if err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(dir, 0750)
		}
	}
	return os.WriteFile(GetPath(fileDir), data, 0644)
}

func JoinAll(fileDir string, f embed.FS) error {
	return JoinPart(fileDir, ".", f)
}

func JoinPart(fileDir, partDir string, f embed.FS) error {
	return fs.WalkDir(f, partDir, func(p string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			fileData, err := f.ReadFile(p)
			if err != nil {
				return err
			}
			trueDir, err := filepath.Rel(partDir, p)
			if err != nil {
				return err
			}
			return Join(filepath.Join(fileDir, trueDir), fileData)
		}
		return nil
	})
}
func Remove(fileDir string) error {
	fileInfo, err := os.Stat(GetPath(fileDir))
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		os.RemoveAll(GetPath(fileDir))
	}
	return nil
}

func GetPath(fileDir string) string {
	return filepath.Join(tempPath, fileDir)
}

func Close() {
	os.RemoveAll(tempPath)
}
