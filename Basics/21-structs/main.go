package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	u1 := User{
		ID:    1,
		Name:  "Huzaifa",
		Email: "huzaifaameer098@gmail.com",
		Age:   22,
	}

	fmt.Println("ID:", u1.ID)
	fmt.Println("Name:", u1.Name)
	fmt.Println("Email:", u1.Email)
	fmt.Println("Age:", u1.Age)
}

//STRUCTS are mutable by default
