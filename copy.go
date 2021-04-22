package copy

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

const (
	homeSymbol = '~'
)

// Copy src to dst
// do nothing if src not exist
func Copy(src, dst string) error {
	src, err := trimHomeSymbol(src)
	if err != nil {
		return fmt.Errorf("failed to trim ~ for src %s", src)
	}

	dst, err = trimHomeSymbol(dst)
	if err != nil {
		return fmt.Errorf("failed to trim ~ for dst %s", dst)
	}

	return copyRaw(src, dst)
}

func copyRaw(src, dst string) error {
	fileInfo, err := os.Stat(src)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return fmt.Errorf("failed to stat src %s: %w", src, err)
	}

	if fileInfo.IsDir() {
		if err := copyDir(src, dst); err != nil {
			return fmt.Errorf("failed to copy dir from src %s to dst %s: %w", src, dst, err)
		}
	} else {
		if err := copyFile(src, dst); err != nil {
			return fmt.Errorf("failed to copy file from src %s to dst %s: %w", src, dst, err)
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
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

func copyDir(src, dst string) error {
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return fmt.Errorf("failed to mkdir %s: %w", dst, err)
	}

	srcFiles, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to read dir %s: %w", src, err)
	}

	for _, srcFile := range srcFiles {
		srcChild := filepath.Join(src, srcFile.Name())
		dstChild := filepath.Join(dst, srcFile.Name())

		if srcFile.IsDir() {
			if err := copyDir(srcChild, dstChild); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcChild, dstChild); err != nil {
				return err
			}
		}
	}

	return nil
}

// trimHomeSymbol replace ~ with full path
// https://stackoverflow.com/a/17609894
func trimHomeSymbol(path string) (string, error) {
	if path == "" || path[0] != homeSymbol {
		return path, nil
	}

	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	newPath := filepath.Join(currentUser.HomeDir, path[1:])
	return newPath, nil
}
