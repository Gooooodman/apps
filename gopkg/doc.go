// Copyright 2014 Gopkg. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Command gopkg helps build packages reproducibly by fixing
their dependencies.

Example Usage

Save currently-used dependencies to file Gopkgs:

	$ gopkg save

Build project using saved dependencies:

	$ gopkg go install

or

	$ GOPATH=`gopkg path`:$GOPATH
	$ go install

*/
package main
