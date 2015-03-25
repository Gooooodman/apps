// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Fix import path.
//
// Example:
//	go run fixpath.go -old=old-path-prefix -new=new-path-prefix
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	flagSourceDir     = flag.String("dir", "", "Set go package dir.")
	flagOldPathPrefix = flag.String("old", "", "Set old import path prefix.")
	flagNewPathPrefix = flag.String("new", "", "Set new import path prefix.")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage of %s: -old=old-path-prefix -new=new-path-prefix\n",
			filepath.Base(os.Args[0]),
		)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if *flagOldPathPrefix == "" || *flagNewPathPrefix == "" {
		flag.Usage()
		return
	}

	// set default dir
	if *flagSourceDir == "" {
		*flagSourceDir = "."
	}

	// import "old-path-prefix...
	if (*flagOldPathPrefix)[0] != '"' {
		*flagOldPathPrefix = `"` + *flagOldPathPrefix
	}
	if (*flagNewPathPrefix)[0] != '"' {
		*flagNewPathPrefix = `"` + *flagNewPathPrefix
	}

	total := 0
	filepath.Walk(*flagSourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("filepath.Walk: ", err)
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		if fixImportPath(path) {
			fmt.Printf("fix %s\n", path)
			total++
		}
		return nil
	})
	fmt.Printf("total %d\n", total)
}

func fixImportPath(path string) bool {
	if path == "fixpath.go" {
		return false // skip self
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("ioutil.ReadFile(%s): ", path, err)
	}
	if !bytes.Contains(data, []byte(*flagOldPathPrefix)) {
		return false
	}
	data = bytes.Replace(data, []byte(*flagOldPathPrefix), []byte(*flagNewPathPrefix), -1)
	if err = ioutil.WriteFile(path, data, 0666); err != nil {
		log.Fatalf("ioutil.WriteFile(%s): ", path, err)
	}
	return true
}
