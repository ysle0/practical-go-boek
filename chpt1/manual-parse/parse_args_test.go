package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestParseArgs(t *testing.T) {
	type testConfig struct {
		args []string
		err  error
		config
	}

	var tests = []testConfig{
		{
			args:   []string{"-h"},
			err:    nil,
			config: config{printUsage: true, numTimes: 0},
		},
		{
			args:   []string{"10"},
			err:    nil,
			config: config{printUsage: false, numTimes: 10},
		},
		{
			args:   []string{"abc"},
			err:    errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
			config: config{printUsage: false, numTimes: 0},
		},
		{
			args:   []string{"1", "foo"},
			err:    errors.New("invalid number of arguments"),
			config: config{printUsage: false, numTimes: 0},
		},
	}
	for i, tc := range tests {
		fmt.Printf("test #%d\n", i)

		c, err := parseArgs(tc.args)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be: %v, got: %v\n",
				tc.err, err)
		}

		if tc.err == nil && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}

		if c.printUsage != tc.config.printUsage {
			t.Errorf("Expected printUsage to be: %v, got: %v\n",
				tc.config.printUsage, c.printUsage)
		}

		if c.numTimes != tc.config.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n",
				tc.config.numTimes, c.numTimes)
		}
	}
}
