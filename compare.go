package copy

import (
	"errors"
	"fmt"
	"os"

	"github.com/pkg/diff"
	"github.com/pkg/diff/write"
)

// Compare src with dst
// return error if src is dir not file
func Compare(src, dst string) error {
	src, err := trimHomeSymbol(src)
	if err != nil {
		return fmt.Errorf("failed to trim ~ for src %s", src)
	}

	dst, err = trimHomeSymbol(dst)
	if err != nil {
		return fmt.Errorf("failed to trim ~ for dst %s", dst)
	}

	return compareRaw(src, dst)
}

func compareRaw(src, dst string) error {
	fileInfo, err := os.Stat(src)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return fmt.Errorf("failed to stat src %s: %w", src, err)
	}

	if fileInfo.IsDir() {
		return fmt.Errorf("currently not support compare dir")
	} else {
		if err := compareFile(src, dst); err != nil {
			return fmt.Errorf("failed to compare file src %s dst %s: %w", src, dst, err)
		}

		return nil
	}
}

func compareFile(src, dst string) error {
	if err := diff.Text(src, dst, nil, nil, os.Stdout, write.TerminalColor()); err != nil {
		return fmt.Errorf("failed to diff text src %s dst %s:%w", src, dst, err)
	}

	return nil
}
