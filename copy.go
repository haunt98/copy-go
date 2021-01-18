package copy

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CopyFile(from, to string) error {
	fromFile, err := os.Open(from)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", from, err)
	}
	defer fromFile.Close()

	toFile, err := os.Create(to)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", to, err)
	}
	defer toFile.Close()

	if _, err := io.Copy(toFile, fromFile); err != nil {
		return fmt.Errorf("failed to copy from %s to %s: %w", from, to, err)
	}

	return nil
}

func CopyDir(from, to string) error {
	if err := os.MkdirAll(to, os.ModePerm); err != nil {
		return fmt.Errorf("failed to mkdir %s: %w", to, err)
	}

	fileInfos, err := ioutil.ReadDir(from)
	if err != nil {
		return fmt.Errorf("failed to read dir %s: %w", from, err)
	}

	for _, fileInfo := range fileInfos {
		newFrom := filepath.Join(from, fileInfo.Name())
		newTo := filepath.Join(to, fileInfo.Name())

		if fileInfo.IsDir() {
			if err := CopyDir(newFrom, newTo); err != nil {
				return err
			}
			continue
		}

		if err := CopyFile(newFrom, newTo); err != nil {
			return err
		}
	}

	return nil
}
