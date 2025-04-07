package main

import "fmt"

func RunArrays() {
	// array in GO is fixed-size, non-changeable
	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//fmt.Println(a[0], a[1])
	//fmt.Println(a)
	//
	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//fmt.Println(primes)

	// slice is dynamically-sized, flexible view into the elements of an array.
	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//
	//s := primes[1:4]
	//fmt.Println(s)

	// slices are like references to arrays.
	// changing the elements of a slice modifies the corresponding elements of its
	// underlying array.

	//names := [4]string{
	//	"John",
	//	"Paul",
	//	"Geroge",
	//	"Ringo",
	//}
	//fmt.Println(names)
	//
	//s1 := names[:2]
	//s2 := names[1:3]
	//fmt.Println(s1, s2)
	//
	//s2[0] = "XXX"
	//fmt.Println("after change s2[0] to \"XXXX\"")
	//fmt.Println("s1: ", s1, "s2: ", s2)
	//fmt.Println("names: ", names)

	//fmt.Println("<< Slice Literals >>")
	//q := []int{2, 3, 5, 7, 11, 13}
	//fmt.Println("q: ", q)
	//
	//r := []bool{true, false, true, true, false, true}
	//fmt.Println(r)
	//
	//s := []struct {
	//	i int
	//	b bool
	//}{
	//	{2, true},
	//	{3, false},
	//	{5, true},
	//	{7, true},
	//	{11, false},
	//	{13, true},
	//}
	//fmt.Println(s)

	// the length of a slice is the number of elements it contains.
	// the capacity of a slice is the number of elements in the underlying array,

	//s := []int{2, 3, 5, 7, 11, 13}
	//printSlice(s)
	//
	//fmt.Println("Extend its length.")
	//s = s[:4]
	//printSlice(s)
	//
	//fmt.Println("drop its first 2 values.")
	//s = s[2:]
	//printSlice(s)

	// you can extend a slice's length by re-slicing it, provided it has sufficient
	// capacity.

	//var s []int
	//fmt.Println(s, len(s), cap(s))
	//if s == nil {
	//	fmt.Println("nil!")
	//}

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	v := make([]int, 0, 5)
	printSlice("v", v)

	v = v[:cap(v)]
	printSlice("v", v)

	v = v[1:]
	printSlice("v", v)
}

func printSlice(str string, s []int) {
	fmt.Printf("%s len=%d cap=%d\n",
		str, len(s), cap(s))
}
