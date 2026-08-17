package main
import (
	"fmt"
	"strings"
)

func main (){
	firstName:= "Huzaifa"
	lastName:= "Ameer"
	fullName:= firstName + "  " + lastName
	fmt.Println(fullName)
	fmt.Println(strings.ToUpper(fullName))
}