package main

import "fmt"

func main() {

	citizensEarnings := make (map[string]float64)

	citizensEarnings["Harrisromeo"] = 35000.0

	citizensEarnings["ijay"] = 29000.0

	citizensEarnings ["Someone"] = 30000.0

	totalTax := task420(citizensEarnings)

	for citizen,tax := range totalTax {

		fmt.Println(citizen, tax)

	}

}

	func task420(citizensEarnings map[string]float64) map[string]float64 {

		totalTax := make(map[string]float64)

		ceilingEarnings := 30000.0

		for citizens, earnings := range citizensEarnings {

			totalEarnings := float64(earnings)

			if totalEarnings <= ceilingEarnings {

				baseTax := 0.15 * totalEarnings

				totalTax[citizens] = baseTax

			} else if totalEarnings > ceilingEarnings {

				baseTax := 0.15 * ceilingEarnings

				additionalIncome := totalEarnings - ceilingEarnings

				additionalTax := 0.20 * additionalIncome

				totalTax[citizens] = baseTax + additionalTax

				}
		}

		return totalTax

	}


