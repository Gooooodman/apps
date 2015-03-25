package fs_test

import (
	"fmt"
	"os"

	"github.com/chai2010/gopkg/apps/gopkg/internal/fs"
)

func ExampleWalker() {
	walker := fs.Walk("/usr/lib")
	for walker.Step() {
		if err := walker.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(walker.Path())
	}
}
