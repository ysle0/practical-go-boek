package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	numTimes int
	name     string
}

var (
	errPosArgsSpecified       = errors.New("position arguments specified")
	errInvalidPosArgSpecified = errors.New("more than one positional argument specified")
)

func main() {
	exitCode := run(os.Args[1:], os.Stdin, os.Stdout)
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func run(args []string, r io.Reader, w io.Writer) int {
	c, err := parseArgs(os.Stderr, args)
	if err != nil {
		if errors.Is(err, errPosArgsSpecified) {
			fmt.Fprintln(w, err)
		}
		return 1
	}

	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(w, err)
		return 1
	}

	err = runCmd(r, w, c)
	if err != nil {
		fmt.Fprintln(w, err)
		return 1
	}

	return 0
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "your name please? Press the Enter key when done.\n"
	io.WriteString(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("you didn't enter your name")
	}
	return name, nil
}

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.Usage = func() {
		var usageString = `
A greeter application which prints the name you entered a specified number of times.
Usage of %s: <options> [name]`
		fmt.Fprintf(w, usageString, fs.Name())
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}
	fs.SetOutput(w)

	fs.IntVar(&c.numTimes, "n", 0, "Number of times to greet")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	if fs.NArg() > 1 {
		return c, errInvalidPosArgSpecified
	}
	if fs.NArg() == 1 {
		c.name = fs.Arg(0)
	}
	return c, nil
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("must specify a number greater than 0")
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	var err error
	if len(c.name) == 0 {
		c.name, err = getName(r, w)
		if err != nil {
			return err
		}
	}
	greetUser(c, w)
	return nil
}

func greetUser(c config, w io.Writer) {
	msg := fmt.Sprintf("leuk u te ontmoeten %s\n", c.name)
	for range c.numTimes {
		io.WriteString(w, msg)
	}
}
