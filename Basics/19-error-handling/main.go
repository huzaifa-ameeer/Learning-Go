// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main (){
// 	if problem := run(); problem != nil {
// 		fmt.Println("Error", problem)
// 	}
// }

// func run () error {
// 	input := "0"
// 	level, err := parseLevel(input)
// 	if err != nil {
// 		return  err
// 	}
// 	fmt.Println("Parsed Level:", level)
// 	return nil
// }

// func parseLevel (s string) (int, error) {
// 	value, err := strconv.Atoi(s)
// 	if err != nil {
// 		return 0, fmt.Errorf("Level must be a number")
// 	}
// 	if value<1 || value>5 {
// 		return 0, fmt.Errorf("Level must be 1 and 5")
// 	}
// 	return value, nil
// }

package main
import "fmt"

func main (){
	if problem := run(); problem != nil {
		fmt.Println("Error:", problem)
	}
}

func run () error {
	input, err := checkValue(6)
	// assign := checkValue(input)
	if err != nil{
		return err
	}
	fmt.Println("Given value:", input)
	return nil
}

func checkValue (value int) (int, error) {

	if value<1 || value>5 {
		return 0, fmt.Errorf("Value must be between 1 and 5")
	}
	return  value, nil
}