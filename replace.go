package copy

import (
	"fmt"
	"os"
)

// Replace dst with src
// delete dst first then copy src to dst
func Replace(src, dst string) error {
	src, err := trimHomeSymbol(src)
	if err != nil {
		return fmt.Errorf("failed to trim ~ src %s", src)
	}

	dst, err = trimHomeSymbol(dst)
	if err != nil {
		return fmt.Errorf("failed to trim ~ dst %s", dst)
	}

	if err := os.RemoveAll(dst); err != nil {
		return fmt.Errorf("failed to remove dst %s: %w", dst, err)
	}

	if err := copyRaw(src, dst); err != nil {
		return fmt.Errorf("failed to copy src %s dst %s: %w", src, dst, err)
	}

	return nil
}
