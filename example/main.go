package main

import (
	"fmt"

	"github.com/haunt98/copy-go"
)

func main() {
	if err := copy.CopyFile("copy.go", "copy.go_cloned"); err != nil {
		fmt.Println(err)
	}

	if err := copy.CopyDir("example", "example_cloned"); err != nil {
		fmt.Println(err)
	}
}
