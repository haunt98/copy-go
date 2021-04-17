package main

import (
	"fmt"

	"github.com/haunt98/copy-go"
)

func main() {
	// Copy file
	if err := copy.Copy("copy.go", "copy.go_cloned"); err != nil {
		fmt.Println(err)
	}

	// Copy dir
	if err := copy.Copy("example", "example_cloned"); err != nil {
		fmt.Println(err)
	}
}
