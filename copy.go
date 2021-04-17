package copy

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type CopyFn func(src, dst string) error

// Copy file from src (source) -> to dst (destination)
// Ignore not exist error
func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("failed to open %s: %w", src, err)
	}
	defer srcFile.Close()

	// Make sure nested dir is exist before copying file
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to mkdir %s: %w", dstDir, err)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", dst, err)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("failed to copy from %s to %s: %w", src, dst, err)
	}

	return nil
}

// Copy dir from src -> to dst
// Ignore not exist error
func CopyDir(src, dst string) error {
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return fmt.Errorf("failed to mkdir %s: %w", dst, err)
	}

	files, err := os.ReadDir(src)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return fmt.Errorf("failed to read dir %s: %w", src, err)
	}

	for _, file := range files {
		tempSrc := filepath.Join(src, file.Name())
		tempDst := filepath.Join(dst, file.Name())

		if file.IsDir() {
			if err := CopyDir(tempSrc, tempDst); err != nil {
				return err
			}

			continue
		}

		if err := CopyFile(tempSrc, tempDst); err != nil {
			return err
		}
	}

	return nil
}
