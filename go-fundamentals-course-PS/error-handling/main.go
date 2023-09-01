package main

import (
	"bufio"
	"demo/go-fundamentals-course-PS/error-handling/menu"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {
loop:
	for {

		fmt.Println("Please select a option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")

		choice, _ := in.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			menu.PrintMenu()
		case "2":
			err := menu.AddItem()
			if err != nil {
				fmt.Println(fmt.Errorf("invalid input: %w", err))
			}
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}
