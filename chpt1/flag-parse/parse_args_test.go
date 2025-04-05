package main

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

// 1. 함수 내에서 새로운 FlagSet 객체를 생성
// 2. FlagSet 객체의 Output() 메서드를 사용하여 FlagSet 의 모든 메서드가 지정된
//    io.Writer 객체의 w 변수로 출력되도록 함
// 3. 파싱할 인수를 매개변수 args 로 전달.

func TestParseArgs(t *testing.T) {
	type testcase struct {
		name     string
		args     []string
		err      error
		numTimes int
	}

	tcs := []testcase{
		{
			name:     "when only help flag is provided",
			args:     []string{"-h"},
			err:      errors.New("flag: help requested"),
			numTimes: 0,
		},
		{
			name:     "when number flag and number of times are provided",
			args:     []string{"-n", "10"},
			err:      nil,
			numTimes: 10,
		},
		{
			name:     "when wrong positional argument is provided with number flag",
			args:     []string{"-n", "abc"},
			err:      errors.New("invalid value \"abc\" for flag -n: parse error"),
			numTimes: 0,
		},
		{
			name:     "when extra wrong positional argument is provided with number flag",
			args:     []string{"-n", "1", "foo"},
			err:      errors.New("positional arguments specified"),
			numTimes: 1,
		},
	}

	byteBuf := new(bytes.Buffer)
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("#%d: %s", i, tc.name), func(t *testing.T) {
			c, err := parseArgs(byteBuf, tc.args)
			if tc.err != nil && err != nil && tc.err.Error() != err.Error() {
				t.Errorf("expected error to be: %v, got: %v\n", tc.err, err)
			}

			if tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("expected error to be: %v, got: %v\n", tc.err, err)
			}

			if c.numTimes != tc.numTimes {
				t.Errorf("expected numTimes to be: %d, got: %d\n", tc.numTimes, c.numTimes)
			}
			byteBuf.Reset()
		})
	}
}
