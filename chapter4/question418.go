package main

import "fmt"

func main() {
	var account int = 1
	var balance, charges, credits, creditLimit, newBal int

	for account != 0 {
		fmt.Print("Input Account Number: ")
		_, err := fmt.Scan(&account)
		if err != nil {
			return
		}

		fmt.Print("Input beginning balance: ")
		_, err = fmt.Scan(&balance)
		if err != nil {
			return
		}

		fmt.Print("Input total charges: ")
		_, err = fmt.Scan(&charges)
		if err != nil {
			return
		}

		fmt.Print("Input total credits: ")
		_, err = fmt.Scan(&credits)
		if err != nil {
			return
		}

		fmt.Print("Input credit limit: ")
		_, err = fmt.Scan(&creditLimit)
		if err != nil {
			return
		}

		newBal = balance + charges - credits
		fmt.Printf("Equivalent New Balnce: %d\n", newBal)
		if newBal > creditLimit {
			fmt.Println("Credit Limit Exceeded")
			break
		}

	}
}
