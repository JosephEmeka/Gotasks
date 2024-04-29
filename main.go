package main

import (
	"fmt"
)

func main() {
	var totalMiles, totalGallons, tripCount int
	var miles, gallons int

	for {
		fmt.Print("Enter miles driven for trip ", tripCount+1, ": ")
		fmt.Scanln(&miles)

		if miles == -1 {
			break
		}

		fmt.Print("Enter gallons used for trip ", tripCount+1, ": ")
		fmt.Scanln(&gallons)

		mpg := float64(miles) / float64(gallons)
		fmt.Printf("Miles per gallon for trip %d: %.2f\n", tripCount+1, mpg)

		totalMiles += miles
		totalGallons += gallons
		totalMpg := float64(totalMiles) / float64(totalGallons)
		fmt.Printf("Combined miles per gallon up to trip %d: %.2f\n", tripCount+1, totalMpg)

		tripCount++
	}

	fmt.Println("Thank you for using the Gas Mileage Calculator.")
}
