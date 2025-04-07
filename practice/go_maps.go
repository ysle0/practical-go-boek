package main

import (
	"fmt"
	"strings"
)

type vertex struct {
	Lat, Long float64
}

var m map[string]vertex

func RunMaps() {
	// make fn on map returns a map of the given type, initialized and ready for use.
	//m = make(map[string]Vertex)
	//
	//m["Bell Labs"] = Vertex{
	//	40.68433,
	//	-74.39967,
	//}
	//fmt.Println(m["Bell Labs"])

	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	if v2, ok2 := m["Answwer"]; ok2 {
		fmt.Println("The value:", v2, "Present?", ok)
	}

	cw := countWords("This is sentences!")
	fmt.Println(cw)
}

func countWords(s string) map[string]int {
	m := make(map[string]int)
	split := strings.Fields(s)

	for _, c := range split {
		if _, ok := m[c]; ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}

	return m
}
