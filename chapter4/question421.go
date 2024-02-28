package main

import "fmt"

func main() {
	var numbers[] int = {1, 3, 2, 7, 6, 8, 9, 3, 5, 1}
	largest := largestOf(numbers)
	fmt.Print(largest)
}
func largestOf(numbers [10]int) int {
	var number int
	for counter := 1; counter <= 10; counter++ {
		numbers[counter], _ = fmt.Scan(&number)
	}
	largest := numbers[0]
	for count := 1; count < len(numbers); count++ {
		if largest < numbers[count] {
			largest = numbers[count]
		}
	}
	return largest
}
