package main

import (
	"fmt"
	"time"
)

func RunErrors() {
	if err := NewMyError(); err != nil {
		fmt.Println(err)
	}

}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("[%v] \n%s\n",
		e.When, e.What)
}

func NewMyError() error {
	return &MyError{
		time.Now(),
		"It did not work",
	}
}
