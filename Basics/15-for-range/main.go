package main

import "fmt"

func main() {
	views := []int{10, 20, 30, 40, 50, 60}
	total := 0

	for i, v := range views {
		fmt.Println("Index:", i, "Views:", v)
		total = total + i
	}

	fmt.Println(total)
}
