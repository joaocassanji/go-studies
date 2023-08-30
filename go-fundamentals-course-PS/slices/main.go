package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s)

	s = []string{"a", "b", "c"}
	fmt.Println(s)

	fmt.Println(s[1])
	s[1] = "x"
	fmt.Println(s[1])

	s2 := s
	s2[2] = "Joao"
	fmt.Println(s, s2)
	s2 = append(s2, "hi")
	fmt.Println(s, s2)

}
