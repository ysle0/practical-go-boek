package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	os.Exit(Run(5 * time.Second))
}

func Run(timeout time.Duration) int {
	cx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	name, err := getNameContext(cx)

	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		return 1
	}
	fmt.Fprintln(os.Stdout, name)
	return 0
}

func getNameContext(cx context.Context) (string, error) {
	var err error
	name := "default name"
	c := make(chan error, 1)
	go func() {
		name, err = getName(os.Stdin, os.Stdout)
		c <- err
	}()

	// wait any of which both
	// 1. timeout-ed and err
	// 2. got name
	select {
	case <-cx.Done():
		return name, cx.Err()
	case err = <-c:
		return name, err
	}
}

func getName(r io.Reader, w io.Writer) (string, error) {
	scan := bufio.NewScanner(r)
	msg := "your name please? press the Enter key when done"
	fmt.Fprintln(w, msg)
	scan.Scan()
	if err := scan.Err(); err != nil {
		return "", err
	}
	name := scan.Text()
	if len(name) == 0 {
		return "", errors.New("you entered an empty name")
	}
	return name, nil
}
