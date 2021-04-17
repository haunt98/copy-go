package copy

import (
	"fmt"
	"os"
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
	replacedSrc, err := trimHomeSymbol(src)
	if err != nil {
		return fmt.Errorf("failed to replace home symbol %s: %w", src, err)
	}

	replacedDst, err := trimHomeSymbol(dst)
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
