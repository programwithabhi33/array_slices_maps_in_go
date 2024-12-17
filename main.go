package main

import "fmt"

func main() {
	prices := [4]float64{11.2, 22.0, 23.3, 12.4}
	fmt.Println("prices are: ", prices)
	slicedPrices := prices[1:3]
	fmt.Println("slicesPrices are: ", slicedPrices)
}
