package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"Small": 1.65, "Medium": 1.80, "Large": 1.95}},
		{name: "Espresso", prices: map[string]float64{"Single": 2.65, "Double": 2.80, "Triple": 2.95}},
	}

	in := bufio.NewReader(os.Stdin)

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
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range item.prices {
					fmt.Printf("\t%10s%10.2f\n", size, price)
				}
			}
		case "2":
			fmt.Println("Please enter the name of the new item")
			name, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name: strings.TrimSpace(name), prices: map[string]float64{"Single": 2.65, "Double": 2.80, "Triple": 2.95}})
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}
