package main
import "fmt"

func main (){
	// points := map[int] int {
	// 	1: 1,
	// 	2: 8,
	// }
	// if val, ok := points[2] ;ok {
	// 	fmt.Println(val)
	// } else {
	// 	fmt.Println("Not found")
	// }

	products := map[string] int {
		"mobile" : 30000,
		"laptop" : 50000,
		"airpods" : 10000,
	}

	total := 0

	for item, price := range products {
		fmt.Println(item, price)
		total = total + price
	}

	fmt.Println("Total is:",total)
}