package main

import "fmt"

func main() {

	i := 0

	for {
		i += 1
		fmt.Println(i)
		break
	}

	i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
