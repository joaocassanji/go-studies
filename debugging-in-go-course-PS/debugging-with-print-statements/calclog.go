package main

import (
	"log"
)

func finalCalcLog(itemOpeningStock int, itemSold int, itemReturned int, itemMissing int) int {

	var itemClosingStock int
	log.Println("itemClosingStock is ", itemClosingStock)
	itemClosingStock = minus(itemOpeningStock, itemSold)
	log.Println("itemSold is ", itemSold)
	log.Println("itemClosingStock is ", itemClosingStock)
	itemClosingStock = plus(itemClosingStock, itemReturned)
	log.Println("itemReturned is ", itemReturned)
	log.Println("itemClosingStock is ", itemClosingStock)
	itemClosingStock = itemClosingStock - itemMissing
	log.Println("itemMissing is ", itemMissing)
	log.Println("itemClosingStock is ", itemClosingStock)

	return (itemClosingStock)
}
