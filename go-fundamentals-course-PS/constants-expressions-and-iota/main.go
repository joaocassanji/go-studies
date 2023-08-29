package main

import "fmt"

func main() {

	const a = 42
	const b float32 = 3
	var f = a
	var g float32 = a
	var f32 float32 = b
	var f64 float64 = float64(b)
	fmt.Println(f, g, f32, f64)

	const c = iota
	fmt.Println(c)

	const (
		g1c1 = 2 * 5
		g1c2 // copy previous value
		g1c3 = iota
		g1c4
		g1c5 = 10 * iota
	)

	fmt.Println(g1c1, g1c2, g1c3, g1c4, g1c5)

}
