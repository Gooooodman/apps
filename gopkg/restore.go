// Copyright 2014 Gopkg. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"path/filepath"
)

var cmdRestore = &Command{
	Usage: "restore",
	Short: "check out listed dependency versions in GOPATH",
	Long: `
Restore checks out the Gopkgs-specified version of each package in GOPATH.
`,
	Run: runRestore,
}

func runRestore(cmd *Command, args []string) {
	g, err := ReadAndLoadGopkgs(findGopkgsJSON())
	if err != nil {
		log.Fatalln(err)
	}
	hadError := false
	for _, dep := range g.Deps {
		err := restore(dep)
		if err != nil {
			log.Println("restore:", err)
			hadError = true
		}
	}
	if hadError {
		os.Exit(1)
	}
}

// restore downloads the given dependency and checks out
// the given revision.
func restore(dep Dependency) error {
	// make sure pkg exists somewhere in GOPATH
	err := runIn(".", "go", "get", "-d", dep.ImportPath)
	if err != nil {
		return err
	}
	ps, err := LoadPackages(dep.ImportPath)
	if err != nil {
		return err
	}
	pkg := ps[0]
	if !dep.vcs.exists(pkg.Dir, dep.Rev) {
		dep.vcs.vcs.Download(pkg.Dir)
	}
	return dep.vcs.RevSync(pkg.Dir, dep.Rev)
}

func findGopkgsJSON() (path string) {
	dir, isDir := findGopkgs()
	if dir == "" {
		log.Fatalln("No Gopkgs found (or in any parent directory)")
	}
	if isDir {
		return filepath.Join(dir, "Gopkgs", "Gopkgs.json")
	}
	return filepath.Join(dir, "Gopkgs")
}
