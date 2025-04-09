package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stdout, "Usage: %s <command> <argument>\n", os.Args[0])
		os.Exit(1)
	}

	cmd, arg := os.Args[1], os.Args[2]
	cmdTimeout := 20 * time.Second
	cx, cancel := createContextWithTimeout(cmdTimeout)
	defer cancel()

	setupSignalHandler(os.Stdout, cancel)

	err := execCommand(cx, cmd, arg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createContextWithTimeout(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d)
}

func setupSignalHandler(w io.Writer, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-ch
		fmt.Fprintf(w, "Received signal: %v\n", sig)
		cancel()
	}()
}

func execCommand(cx context.Context, cmd string, arg string) error {
	return exec.CommandContext(cx, cmd, arg).Run()
}
