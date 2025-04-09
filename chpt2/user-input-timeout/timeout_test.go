package main

import (
	"os"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	t.Run("OK: timeout should work", func(t *testing.T) {
		exitCode := Run(5 * time.Millisecond)
		if exitCode == 0 {
			t.Fatalf("expect: exit code == 1, actual: %d", exitCode)
		}
	})

	t.Run("NO: should not timeout", func(t *testing.T) {
		// Run the test with a short timeout
		// Create a pipe
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		// Save original stdin and restore it later
		originalStdin := os.Stdin
		os.Stdin = r
		defer func() {
			os.Stdin = originalStdin
			r.Close()
			w.Close()
		}()

		// Write input to the pipe
		go func() {
			_, err := w.WriteString("test input\n")
			if err != nil {
				t.Error(err)
			}
		}()

		exitCode := Run(10 * time.Millisecond)
		if exitCode != 0 {
			t.Fatalf("expect: exit code == 0, actual: %d", exitCode)
		}
	})
}
