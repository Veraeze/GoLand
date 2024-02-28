package main

import "fmt"

func main() {
	var mileage, gallon int
	var totalMileage int = 0
	var totalGallon int = 0
	var counter int = 1

	fmt.Print("Enter miles driven (or -1 to quit): ")
	_, err := fmt.Scan(&mileage)
	if err != nil {
		return
	}
	fmt.Print("Enter gallons used (or -1 to quit): ")
	_, err = fmt.Scan(&gallon)
	if err != nil {
		return
	}

	average := float64(mileage / gallon)
	totalAverage := average
	fmt.Printf("The miles per gallons used for this trip is: %.2f\n", average)

	for mileage != -1 && gallon != -1 {

		totalMileage = totalMileage + mileage
		totalGallon = totalGallon + gallon
		counter = counter + 1

		fmt.Print("Enter miles driven (or -1 to quit): ")
		_, err := fmt.Scan(&mileage)
		if err != nil {
			return
		}
		fmt.Print("Enter gallons used (or -1 to quit): ")
		_, err = fmt.Scan(&gallon)
		if err != nil {
			return
		}
		if mileage == -1 && gallon == -1 {
			break
		}

		average = float64(mileage / gallon)
		totalAverage = totalAverage + average
		fmt.Printf("The miles per gallons used for this trip is: %.2f\n", average)
	}

	fmt.Printf("The combined miles per gallons used for all the trips is: %.2f\n", totalAverage)

}
