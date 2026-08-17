package main

import "fmt"

func main() {
	// scores := make([]int, 0, 5)

	// scores = append(scores, 10, 20, 30, 40, 50, 60)

	// fmt.Println(scores)
	// fmt.Println("Length of slice:", len(scores))
	// fmt.Println("Capacity of slice:", cap(scores))

	todos := [] string {"Learn golang", "Make some MERN stack projects"}
	more := [] string {"Lean Nest.js"}

	todos = append(todos, more...)
	fmt.Println(todos)
}
