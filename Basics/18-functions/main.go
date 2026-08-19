package main
import "fmt"

func add (a int, b int) int {
	return a+b
}

func addAndProduct (a int, b int) (int, int) {
	sum := a+b
	product := a*b
	return sum, product
}

func main (){
	sum := add(2,4)
	// fmt.Println(sum)
	result1, result2 := addAndProduct(3,4)
	fmt.Println("Sum:", sum, "Results: ",result1,",", result2)
}