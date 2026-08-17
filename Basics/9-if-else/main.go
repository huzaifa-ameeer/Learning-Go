package main

import "fmt"

func main() {
	score := 85

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

	items := 5
	pricePerItem := 9

	if totalPrice := items * pricePerItem; totalPrice >=50{
		fmt.Println("Eligible")
	} else {
		fmt.Println("Not Eligible")
	}

}
