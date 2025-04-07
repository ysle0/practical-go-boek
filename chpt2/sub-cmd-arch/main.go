package main

import (
	"errors"
	"fmt"
	"github.com/ysle0/chpt2/sub-cmd-arch/cmd"
	"io"
	"os"
)

func main() {
	err := handleCommand(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}

var (
	errInvalidSubCmd = errors.New("invalid sub command specified")
)

func showUsage(w io.Writer) {
	fmt.Fprintf(w, "usage: mync [http|grpc] -h\n")
	cmd.HandleHttp(w, []string{"-h"})
	cmd.HandleGrpc(w, []string{"-h"})
}

func handleCommand(w io.Writer, args []string) error {
	var err error
	if len(args) < 1 {
		err = errInvalidSubCmd
	} else {
		switch args[0] {
		case "http":
			err = cmd.HandleHttp(w, args[1:])
		case "grpc":
			err = cmd.HandleGrpc(w, args[1:])
		case "-h":
			fallthrough
		case "--help":
			showUsage(w)
		default:
			err = errInvalidSubCmd
		}
	}

	isInvalidSubCmd := errors.Is(err, errInvalidSubCmd)
	isNoServerSpecified := errors.Is(err, cmd.ErrNoServerSpecified)
	if isInvalidSubCmd || isNoServerSpecified {
		fmt.Fprintln(w, err)
		showUsage(w)
	}
	return err
}
