package copy

import (
	"bytes"
	"fmt"

	"github.com/pkg/diff"
)

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
	return nil
}

func compareFile(src, dst string) (string, error) {
	buf := new(bytes.Buffer)

	if err := diff.Text(src, dst, nil, nil, buf); err != nil {
		return "", fmt.Errorf("failed to diff text src %s dst %s:%w", src, dst, err)
	}

	return buf.String(), nil
}
