package copy

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const (
	homeSymbol = '~'
)

// ReplaceFile remove file in dst and replace with file in src
func ReplaceFile(src, dst string) error {
	return replace(src, dst, CopyFile)
}

// ReplaceDir remove dir in dst and replace with dir src
func ReplaceDir(src, dst string) error {
	return replace(src, dst, CopyDir)
}

// replace dst with src
func replace(src, dst string, copyFn CopyFn) error {
	replacedSrc, err := replaceHomeSymbol(src)
	if err != nil {
		return fmt.Errorf("failed to replace home symbol %s: %w", src, err)
	}

	replacedDst, err := replaceHomeSymbol(dst)
	if err != nil {
		return fmt.Errorf("failed to replace home symbol %s: %w", dst, err)
	}

	if err := os.RemoveAll(replacedDst); err != nil {
		return fmt.Errorf("failed to remove %s: %w", replacedDst, err)
	}

	if err := copyFn(replacedSrc, replacedDst); err != nil {
		return fmt.Errorf("failed to copy from %s to %s: %w", replacedSrc, replacedDst, err)
	}

	return nil
}

// replaceHomeSymbol replace ~ with full path
// https://stackoverflow.com/a/17609894
func replaceHomeSymbol(path string) (string, error) {
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
