package copy

import (
	"bytes"
	"fmt"

	"github.com/pkg/diff"
)

func Compare(src, dst string) error {
	return nil
}

func compareFile(src, dst string) (string, error) {
	buf := new(bytes.Buffer)

	if err := diff.Text(src, dst, nil, nil, buf); err != nil {
		return "", fmt.Errorf("failed to diff text src %s dst %s:%w", src, dst, err)
	}

	return buf.String(), nil
}
