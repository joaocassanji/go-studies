package main

import "fmt"

func main() {
	var arr [3]string
	fmt.Println(arr)

	arr = [3]string{"a", "b", "c"}
	fmt.Println(arr)

	fmt.Println(arr[1])
	arr[1] = "x"
	fmt.Println(arr[1])

	arr2 := arr
	arr2[2] = "Joao"
	fmt.Println(arr, arr2)

}
