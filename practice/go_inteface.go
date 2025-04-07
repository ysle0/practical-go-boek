package main

import (
	"fmt"
	"math"
)

func RunInterface() {
	//runAbser()

	//var i interface{} = "hello"
	//s := i.(string)
	//fmt.Println(s)
	//
	//s, ok := i.(string)
	//fmt.Println(s, ok)
	//
	//f, ok := i.(float64)
	//fmt.Println(f, ok)
	//
	//f = i.(float64)
	//fmt.Println(f)

	//do(21)
	//do("hello")
	//do(true)

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblerox", 9001}
	fmt.Println(a, z)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v jaren)", p.Name, p.Age)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func runAbser() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	//a = v
	fmt.Println(a.Abs())
}

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (x MyFloat) Abs() float64 {
	if x < 0 {
		return float64(-x)
	}
	return float64(x)
}

type Vertex struct {
	X, Y float64
}

func (x *Vertex) Abs() float64 {
	return math.Sqrt(x.X*x.X + x.Y*x.Y)
}
