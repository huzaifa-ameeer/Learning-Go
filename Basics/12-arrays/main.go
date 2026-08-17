package main
import "fmt"

func main (){

	//fixed size
	var array [3] int
	array[0] = 10
	array[1] = 20
	array[2] = 30

	fmt.Println("Array 1:", array)

	//aray literals

	array1 := [3] int {1,5,9}
	fmt.Println("Array 2:", array1)
}