// package main
// import "fmt"

// func add (a int, b int) int {
// 	return a+b
// }

// func addAndProduct (a int, b int) (int, int) {
// 	sum := a+b
// 	product := a*b
// 	return sum, product
// }

// func main (){
// 	sum := add(2,4)
// 	// fmt.Println(sum)
// 	result1, result2 := addAndProduct(3,4)
// 	fmt.Println("Sum:", sum, "Results: ",result1,",", result2)
// }

// package main
// import "fmt"

// func addAndSub (a int, b int) (sum int, sub int){
// 	sum = a+b
// 	sub = a-b
// 	return
// }

// func main (){
// 	r1, r2 := addAndSub(2,3)
// 	fmt.Println("Sum:", r1, "Sub:", r2)
// }

// package main 
// import "fmt"

// func sumAll (nums... int) int {
// 	total := 0
// 	for _, value := range nums{
// 		total = total + value
// 	}
// 	return total
// }

// func main (){
// 	// fmt.Println(sumAll(1,2,3,4,5,6,7,8,9,10))

// 	values := [] int {1,2,3,4,5}
// 	fmt.Println(sumAll(values...))
// }

package main
import "fmt"

func main (){
	res := func(a int, b int) int {
		return a + b
	} (5,6)

	fmt.Println(res)
}