package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestParseArgs(t *testing.T) {
	type testcase struct {
		config
		args   []string
		output string
		err    error
	}

	tcs := []testcase{
		{
			args:   []string{"-h"},
			config: config{numTimes: 0},
			err:    errors.New("flag: help requested"),
			output: `
A greeter application which prints the name you entered a specified number of times.
Usage of greeter: <options> [name]
Options:
  -n int
	    Number of times to greet
`,
		},
		{
			args:   []string{"-n", "10"},
			config: config{numTimes: 10},
		},
		{
			args:   []string{"-n", "abc"},
			config: config{numTimes: 0},
			err:    errors.New("invalid value \"abc\" for flag -n: parse error"),
		},
		{
			args:   []string{"-n", "1", "John Doe"},
			config: config{numTimes: 1, name: "John Doe"},
		},
		{
			args:   []string{"-n", "1", "John", "Doe"},
			err:    errors.New("more than one positional argument specified"),
			config: config{numTimes: 1},
		},
	}

	byteBuf := new(bytes.Buffer)
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("#%d: %s", i, tc.name), func(t *testing.T) {
			c, err := parseArgs(byteBuf, tc.args)
			if tc.err != nil && err != nil && err.Error() != tc.err.Error() {
				t.Fatalf("expected nil error, got: %v\n", err)
			}
			if tc.err != nil && err.Error() != tc.err.Error() {
				t.Fatalf("expected error to be: %v, got: %v\n", tc.err, err)
			}
			if c.numTimes != tc.config.numTimes {
				t.Fatalf("expected number of times to be: %v, got: %v\n",
					tc.config.numTimes, c.numTimes)
			}
			gotMsg := byteBuf.String()
			tc.output = strings.ReplaceAll(tc.output, "\t", "")
			gotMsg = strings.ReplaceAll(gotMsg, "\t", "")

			if len(tc.output) != 0 && gotMsg != tc.output {
				t.Fatalf("expected output to be: %v, got: %v\n",
					tc.output, gotMsg)
			}
			byteBuf.Reset()
		})
	}
}
