package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidateArgs(t *testing.T) {
	var tests = []struct {
		c   config
		err error
	}{
		{
			c:   config{},
			err: errors.New("must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: -1},
			err: errors.New("must specify a number greater than 0"),
		},
		{
			c:   config{numTimes: 10},
			err: nil,
		},
	}

	for i, tc := range tests {
		fmt.Printf("test #%d\n", i)

		err := validateArgs(tc.c)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Errorf("expected error to be %v, but got %v\n", tc.err, err)
		}

		if tc.err == nil && err != nil {
			t.Errorf("expected nil error, got: %v\n", err)
		}
	}
}
