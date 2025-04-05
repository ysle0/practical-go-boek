package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type config struct {
	numTimes   int
	printUsage bool
}

var usageStr = fmt.Sprintf(`Usage: %s <integer> [-h|--help]

A greeter application which prints the name you entered <integer> number of times.`, os.Args[0])

func main() {
	println("starting app\n")
	err := run(os.Args[1:], os.Stdin, os.Stdout)
	if err != nil {
		os.Exit(1)
	}
}

func run(args []string, r io.Reader, w io.Writer) error {
	c, err := parseArgs(args)
	if err != nil {
		fmt.Fprintln(w, err)
		printUsage(w)
		return err
	}

	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(w, err)
		printUsage(w)
		return err
	}

	err = runCmd(r, w, c)
	if err != nil {
		fmt.Fprintln(w, err)
		return err
	}

	return nil
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"
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

func parseArgs(args []string) (config, error) {
	var numTimes int
	var err error
	c := config{}
	if len(args) != 1 {
		return c, errors.New("invalid number of arguments")
	}

	if args[0] == "-h" || args[0] == "--help" {
		c.printUsage = true
		return c, nil
	}

	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	c.numTimes = numTimes

	return c, nil
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("must specify a number greater than 0")
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}

	name, err := getName(r, w)
	if err != nil {
		return err
	}

	greetUser(c, name, w)
	return nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("leuk u te ontmoeten %s\n", name)
	for range c.numTimes {
		io.WriteString(w, msg)
	}
}

func printUsage(w io.Writer) {
	io.WriteString(w, usageStr)
}
