package slice

import "fmt"

func PlayWithSlices() {
	prices := []float64{11.2, 22.0, 23.3, 12.4}
	fmt.Println("The prices is: ", prices)
	updatedSlice := append(prices, 7.99) //Append element on the slice but never update the original slice based on underlying array, it's actually return the brand new slice containing all the element on the slice or array passed as the first argument to append fn and appended element
	fmt.Println("Updated slice is: ", updatedSlice)
	fmt.Println("After append the prices slice is: ", prices)
	// Want to remove elements from array or slice create a slice from the array or slice
}
