package cmd

import "fmt"

var (
	ErrNoServerSpecified = fmt.Errorf("you have to specify the remote server.")
)
