package main

import "fmt"

func main() {
	testScores := map[string]float64{
		"Joao0": 87.3,
		"Joao1": 105,
		"Joao2": 63.5,
		"Joao3": 27,
	}

	c := clone(testScores)
	fmt.Println(c)

}

func clone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}
