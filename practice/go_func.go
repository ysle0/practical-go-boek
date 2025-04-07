package main

func RunFunc() {
	//hybot := func(x, y float64) float64 {
	//	return math.Sqrt(x*x + y*y)
	//}
	//fmt.Println(hybot(5, 12))
	//
	//fmt.Println(compute(hybot))
	//fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 5)
}

func fibonacci() func() int {
	n1, n2, n3 := 0, 1, 1
	return func() int {
		defer func() {
			next := n2 + n3
			n1 = n2
			n2 = n3
			n3 = next
		}()
		return n1
	}
}
