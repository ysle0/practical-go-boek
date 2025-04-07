package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	if len(os.Args) < 2 {
		printUsage(os.Stdout)
		os.Exit(1)
	}

	w := os.Stdout
	switch os.Args[1] {
	case "cmd-a":
		err = handleCmdA(w, os.Args[2:])
	case "cmd-b":
		err = handleCmdB(w, os.Args[2:])
	default:
		printUsage(w)
	}

	if err != nil {
		fmt.Println(err)
	}
	os.Exit(1)
}

func handleCmdA(w io.Writer, args []string) error {
	var v string
	fs := flag.NewFlagSet("cmd-b", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "argument-value", "Argument 1")
	err := fs.Parse(args)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Executing command A")
	return nil
}

func handleCmdB(w io.Writer, s []string) error {
	var v string
	fs := flag.NewFlagSet("cmd-b", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "argument-value", "Argument 2")
	err := fs.Parse(s)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Executing command B")
	return nil
}

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: %s [cmd-a|cmd-b] -h\n", os.Args[0])
	handleCmdA(w, []string{"-h"})
	handleCmdB(w, []string{"-h"})
}
