// Copyright 2014 Gopkg. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"os/exec"
)

var cmdGet = &Command{
	Usage: "get [packages]",
	Short: "download and install packages with specified dependencies",
	Long: `
Get downloads to GOPATH the packages named by the import paths, and installs
them with the dependencies specified in their Gopkgs files.

If any of the packages do not have Gopkgs files, those are installed
as if by go get.

For more about specifying packages, see 'go help packages'.
`,
	Run: runGet,
}

func runGet(cmd *Command, args []string) {
	if len(args) == 0 {
		args = []string{"."}
	}

	err := command("go", "get", "-d", args).Run()
	if err != nil {
		log.Fatalln(err)
	}

	// group import paths by Gopkgs location
	groups := make(map[string][]string)
	ps, err := LoadPackages(args...)
	if err != nil {
		log.Fatalln(err)
	}
	for _, pkg := range ps {
		if pkg.Error.Err != "" {
			log.Fatalln(pkg.Error.Err)
		}
		dir, _ := findInParents(pkg.Dir, "Gopkgs")
		groups[dir] = append(groups[dir], pkg.ImportPath)
	}
	for dir, packages := range groups {
		var c *exec.Cmd
		if dir == "" {
			c = command("go", "install", packages)
		} else {
			c = command("gopkg", "go", "install", packages)
			c.Dir = dir
		}
		if err := c.Run(); err != nil {
			log.Fatalln(err)
		}
	}
}

// command is like exec.Command, but the returned
// Cmd inherits stderr from the current process, and
// elements of args may be either string or []string.
func command(name string, args ...interface{}) *exec.Cmd {
	var a []string
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			a = append(a, v)
		case []string:
			a = append(a, v...)
		}
	}
	c := exec.Command(name, a...)
	c.Stderr = os.Stderr
	return c
}
