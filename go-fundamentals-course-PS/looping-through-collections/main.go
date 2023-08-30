package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please select a option")
	fmt.Println("1) Print menu")
	in := bufio.NewReader(os.Stdin)
	choise, _ := in.ReadString('\n')
	choise = strings.TrimSpace(choise)

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"Small": 1.65, "Medium": 1.80, "Large": 1.95}},
		{name: "Espresso", prices: map[string]float64{"Single": 2.65, "Double": 2.80, "Triple": 2.95}},
	}

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}
