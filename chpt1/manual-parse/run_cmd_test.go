package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		c      config
		input  string
		output string
		err    error
	}{
		{
			c:      config{printUsage: true},
			output: usageStr,
		},
		{
			c:      config{numTimes: 5},
			input:  "",
			output: strings.Repeat("Your name please? Press the Enter key when done.\n", 1),
			err:    errors.New("you didn't enter your name"),
		},
		{
			c:     config{numTimes: 5},
			input: "Bill Bryson",
			output: "Your name please? Press the Enter key when done.\n" +
				strings.Repeat("Leuk u te ontmoeten Bill Bryson\n", 5),
		},
	}

	buf := new(bytes.Buffer)
	for i, tc := range tests {
		fmt.Printf("test #%d\n", i)

		rd := strings.NewReader(tc.input)
		err := runCmd(rd, buf, tc.c)
		if err != nil && tc.err == nil {
			t.Fatalf("expected nil error, got %v", err)
		}

		if tc.err != nil {
			if err.Error() != tc.err.Error() {
				t.Fatalf("expected error to be %v, but got %v", tc.err.Error(), err.Error())
			}
		}

		gotMsg := buf.String()
		if gotMsg != tc.output {
			t.Errorf("expected stdout message to be %v, but got %v", tc.output, gotMsg)
		}
		buf.Reset()
	}
}
