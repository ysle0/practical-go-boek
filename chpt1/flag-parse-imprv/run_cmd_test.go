package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	type testcase struct {
		c      config
		input  string
		output string
		err    error
	}

	tcs := []testcase{
		{
			c:      config{numTimes: 5},
			input:  "",
			output: strings.Repeat("your name please? Press the Enter key when done.\n", 1),
			err:    errors.New("you didn't enter your name"),
		},
		{
			c:     config{numTimes: 5},
			input: "Bill Bryson",
			output: "your name please? Press the Enter key when done.\n" +
				strings.Repeat("leuk u te ontmoeten Bill Bryson\n", 5),
		},
	}

	buf := new(bytes.Buffer)
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			rd := strings.NewReader(tc.input)
			err := runCmd(rd, buf, tc.c)
			if err != nil && tc.err == nil {
				t.Fatalf("expected nil error, got: %v\n", err)
			}
			if tc.err != nil {
				if err.Error() != tc.err.Error() {
					t.Fatalf("expected error to be: %v, got: %v\n", tc.err, err)
				}
			}
			gotMsg := buf.String()
			if gotMsg != tc.output {
				t.Errorf("expected stdout message to be\n %v, but got\n %v", tc.output, gotMsg)
			}
			buf.Reset()
		})
	}
}
