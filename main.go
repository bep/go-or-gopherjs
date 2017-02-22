// Copyright © 2017-present Bjørn Erik Pedersen <bjorn.erik.pedersen@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	goCommand := "go"

	if os.Getenv("GOARCH") == "js" {
		goCommand = "gopherjs"
		// TODO(bep) See https://github.com/bep/go-or-gopherjs/issues/1
		os.Setenv("GOARCH", "")
	}

	cmd := exec.Command(goCommand, os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok {
		// TODO(bep) Windows, see https://groups.google.com/forum/#!msg/golang-nuts/8XIlxWgpdJw/Z8s2N-SoWHsJ
		if status, ok := e.Sys().(syscall.WaitStatus); ok {
			os.Exit(status.ExitStatus())
		}
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}
}
