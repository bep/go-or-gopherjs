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

	// TODO(bep) I thought I should be clever and use the existing GOARCH
	// env variable, but that seems to fail with gopherjs test
	// I will create an issue on the gopherjs issue tracker.
	if os.Getenv("MYGOARCH") == "js" {
		goCommand = "gopherjs"
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
