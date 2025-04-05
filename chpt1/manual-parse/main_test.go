package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	type testFixture struct {
		name  string
		args  []string
		input string
	}

	fails := []testFixture{
		{
			name:  "help flag",
			args:  []string{"-h"},
			input: "",
		},
		{
			name:  "help flag with number",
			args:  []string{"-h", "10"},
			input: "",
		},
		{
			name:  "help flag with number and foo",
			args:  []string{"-h", "10", "foo"},
			input: "",
		},
		{
			name:  "no args",
			args:  []string{},
			input: "",
		},
		{
			name:  "invalid number",
			args:  []string{"abc"},
			input: "",
		},
		{
			name:  "negative number",
			args:  []string{"-1"},
			input: "",
		},
		{
			name:  "zero number",
			args:  []string{"0"},
			input: "",
		},
	}

	succeeds := []testFixture{
		{
			name:  "correct",
			args:  []string{"1"},
			input: "Bill Bryson",
		},
	}

	for _, tc := range fails {
		t.Run(tc.name, func(t *testing.T) {
			w, r := new(bytes.Buffer), strings.NewReader(tc.input)

			err := run(tc.args, r, w)
			if err == nil {
				t.Errorf("error have to be thrown --> %v\n", err)
			}
		})
	}

	for _, tc := range succeeds {
		t.Run(tc.name, func(t *testing.T) {
			w, r := new(bytes.Buffer), strings.NewReader(tc.input)

			err := run(tc.args, r, w)
			if err != nil {
				t.Errorf("error mustn't be thrown! --> %v\n", err)
			}
		})
	}

}
