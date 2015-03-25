// Copyright 2014 Gopkg. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

// A Command is an implementation of a gopkg command
// like gopkg save or gopkg go.
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string)

	// Usage is the one-line usage message.
	// The first word in the line is taken to be the command name.
	Usage string

	// Short is the short description shown in the 'gopkg help' output.
	Short string

	// Long is the long message shown in the
	// 'gopkg help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet
}

func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) UsageExit() {
	fmt.Fprintf(os.Stderr, "Usage: gopkg %s\n\n", c.Usage)
	fmt.Fprintf(os.Stderr, "Run 'gopkg help %s' for help.\n", c.Name())
	os.Exit(2)
}

// Commands lists the available commands and help topics.
// The order here is the order in which they are printed
// by 'gopkg help'.
var commands = []*Command{
	cmdSave,
	cmdGo,
	cmdGet,
	cmdPath,
	cmdRestore,
	cmdUpdate,
}

func main() {
	flag.Usage = usageExit
	flag.Parse()
	log.SetFlags(0)
	log.SetPrefix("gopkg: ")
	args := flag.Args()
	if len(args) < 1 {
		usageExit()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			cmd.Flag.Usage = func() { cmd.UsageExit() }
			cmd.Flag.Parse(args[1:])
			cmd.Run(cmd, cmd.Flag.Args())
			return
		}
	}

	fmt.Fprintf(os.Stderr, "gopkg: unknown command %q\n", args[0])
	fmt.Fprintf(os.Stderr, "Run 'gopkg help' for usage.\n")
	os.Exit(2)
}

var usageTemplate = `
Gopkg is a tool for managing Go package dependencies.

Usage:

    gopkg command [arguments]

The commands are:
{{range .}}
    {{.Name | printf "%-8s"}} {{.Short}}{{end}}

Use "gopkg help [command]" for more information about a command.
`

var helpTemplate = `
Usage: gopkg {{.Usage}}

{{.Long | trim}}
`

func help(args []string) {
	if len(args) == 0 {
		printUsage(os.Stdout)
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: gopkg help command\n\n")
		fmt.Fprintf(os.Stderr, "Too many arguments given.\n")
		os.Exit(2)
	}
	for _, cmd := range commands {
		if cmd.Name() == args[0] {
			tmpl(os.Stdout, helpTemplate, cmd)
			return
		}
	}
}

func usageExit() {
	printUsage(os.Stderr)
	os.Exit(2)
}

func printUsage(w io.Writer) {
	tmpl(w, usageTemplate, commands)
}

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{
		"trim": strings.TrimSpace,
	})
	template.Must(t.Parse(strings.TrimSpace(text) + "\n\n"))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}
