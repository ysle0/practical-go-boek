package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
)

type config struct {
	numTimes       int
	pathToSaveHtml string
}

func main() {
	println("starting app\n")
	err := run(os.Args[1:], os.Stdin, os.Stdout)
	if err != nil {
		os.Exit(1)
	}
}

func run(args []string, r io.Reader, w io.Writer) error {
	c, err := parseArgs(os.Stderr, args)
	if err != nil {
		fmt.Fprintln(w, err)
		return err
	}

	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(w, err)
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
	fset := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fset.SetOutput(w)

	fset.IntVar(&c.numTimes, "n", 0, "Number of times to greet")
	fset.StringVar(&c.pathToSaveHtml, "o", "", "Path to save the html file")
	err := fset.Parse(args)
	if err != nil {
		return c, err
	}
	if fset.NArg() != 0 {
		return c, errors.New("positional arguments specified")
	}
	return c, nil
}

func makeGreeterHtmlPage(c config, name string) error {
	var tmpl *template.Template
	var err error
	tmplContent := `
		{{define "Name"}}
		<!DOCTYPE html>
		<html>
		<head>
			<title>Greeter</title>
		</head>
		<body>
			<h1>Greeter</h1>
				<h1>Hello, {{.Name}}!</h1>
		</body>
		</html>
		{{end}}
	`
	tmpl, err = template.New("greeter").Parse(tmplContent)
	if err != nil {
		fmt.Errorf("error creating template: %v", err)
		return err
	}

	filew := new(bytes.Buffer)
	tmplVarName := "Name"
	err = tmpl.ExecuteTemplate(filew, tmplVarName, map[string]string{tmplVarName: name})
	if err != nil {
		fmt.Errorf("error executing template: %v", err)
		return err
	}

	err = os.WriteFile(c.pathToSaveHtml, filew.Bytes(), 0666)
	if err != nil {
		fmt.Errorf("error writing file: %v", err)
		return err
	}
	return nil
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("must specify a number greater than 0")
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	name, err := getName(r, w)
	if err != nil {
		return err
	}

	if c.pathToSaveHtml != "" {
		err = makeGreeterHtmlPage(c, name)
		if err != nil {
			return err
		}
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
