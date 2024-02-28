package main



import "fmt"

func main() {

	items := make(map[int]float64)

	items[1] = 239.99

	items[2] = 129.75

	items[3] = 99.95

	items[4] = 350.85

	earnings := task419(items)

	for i := 1; i <= len(earnings); i++ {

		fmt.Println(earnings[1])
	}
}

	func task419(items map[int ]float64) map[int ]float64 {

		earnings := make(map [int]float64)

		for i := 1; i <= len(items); i++ {

			grossSale := items[1]

			earning := 200 + (0.09 * grossSale)

			earnings[1] = earning
		}

		return earnings

}



